package db_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/denisK-H/task-6/internal/db"
)

func TestGetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка создания мока: %s", err)
	}
	defer db.Close()

	service := New(db)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob"))

	names, err := service.GetNames()
	assert.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(fmt.Errorf("db error"))

	_, err = service.GetNames()
	assert.ErrorContains(t, err, "db query: db error")

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	_, err = service.GetNames()
	assert.ErrorContains(t, err, "rows scanning:")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(1, fmt.Errorf("scan error"))
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	_, err = service.GetNames()
	assert.ErrorContains(t, err, "rows error: scan error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не все ожидания выполнены: %s", err)
	}
}

func TestGetUniqueNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка создания мока: %s", err)
	}
	defer db.Close()

	service := New(db)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob"))

	names, err := service.GetUniqueNames()
	assert.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(fmt.Errorf("db error"))

	_, err = service.GetUniqueNames()
	assert.ErrorContains(t, err, "db query: db error")

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	assert.ErrorContains(t, err, "rows scanning:")

	rows = sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(1, fmt.Errorf("scan error"))
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	assert.ErrorContains(t, err, "rows error: scan error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не все ожидания выполнены: %s", err)
	}
}
