package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/denisK-H/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var (
	errDB   = errors.New("db error")
	errScan = errors.New("scan error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob"))

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(errDB)

	_, err = service.GetNames()
	require.ErrorContains(t, err, "db query: db error")

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	_, err = service.GetNames()
	require.ErrorContains(t, err, "rows scanning:")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(1, errScan)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	_, err = service.GetNames()
	require.ErrorContains(t, err, "rows error: scan error")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob").CloseError(errScan)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	_, err = service.GetNames()
	require.ErrorContains(t, err, "rows error: scan error")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob"))

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errDB)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, "db query: db error")

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, "rows scanning:")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(1, errScan)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, "rows error: scan error")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob").CloseError(errScan)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, "rows error: scan error")

	require.NoError(t, mock.ExpectationsWereMet())
}
