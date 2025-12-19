package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VlasfimosY/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var errClose = errors.New("close error")

func TestGetNames(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice"))

	dbService := db.New(dbConn)
	names, err := dbService.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	dbService := db.New(dbConn)
	names, err := dbService.GetNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrNoRows)

	dbService := db.New(dbConn)
	_, err = dbService.GetNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(dbConn)
	_, err = dbService.GetNames()
	require.ErrorContains(t, err, "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Bob"))

	dbService := db.New(dbConn)
	names, err := dbService.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Bob"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Empty(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	dbService := db.New(dbConn)
	names, err := dbService.GetUniqueNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrNoRows)

	dbService := db.New(dbConn)
	_, err = dbService.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(dbConn)
	_, err = dbService.GetUniqueNames()
	require.ErrorContains(t, err, "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErr(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").CloseError(errClose)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(dbConn)
	_, err = dbService.GetNames()
	require.ErrorContains(t, err, "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErr(t *testing.T) {
	t.Parallel()

	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Bob").CloseError(errClose)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(dbConn)
	_, err = dbService.GetUniqueNames()
	require.ErrorContains(t, err, "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}
