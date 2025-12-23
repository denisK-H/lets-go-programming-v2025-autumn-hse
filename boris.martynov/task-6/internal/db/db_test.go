package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JingolBong/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

const (
	nameQuery        = "SELECT name FROM users"
	uniqueNamesQuery = "SELECT DISTINCT name FROM users"
)

var (
	errQuery    = errors.New("db query: ")
	errRowScann = errors.New("rows scanning: ")
	errRow      = errors.New("rows error: ")
	errClose    = errors.New("rows close")
)

func TestDBGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).AddRow("Jingol")

	mock.ExpectQuery(nameQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Jingol"}, name)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBGetNamesErrorQuery(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery(nameQuery).WillReturnError(errQuery)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errQuery.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBGetNamesErrorScan(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(nameQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errRowScann.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBGetNamesErrorRow(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).CloseError(errClose)
	mock.ExpectQuery(nameQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errRow.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBUniqueGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).AddRow("Jingol")

	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Jingol"}, name)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBUniqueGetNamesErrorQuery(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery(uniqueNamesQuery).WillReturnError(errQuery)

	service := db.New(mockDB)
	name, err := service.GetUniqueNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errQuery.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBUniqueGetNamesErrorScan(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetUniqueNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errRowScann.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBUniqueGetNamesErrorRow(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).CloseError(errClose)
	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetUniqueNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errRow.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
