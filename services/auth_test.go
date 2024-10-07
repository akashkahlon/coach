package services

import (
	"coach/testutils"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	testutils.SetupTestEnv(m)
	testutils.CleanUpDatabase()
}



func TestLoginUser(t *testing.T) {
	email := "random@y.com"
	password := "password"
	testutils.CreateUser(email, password)

	token, tokenExpirationTime, err := LoginUser(testutils.TestDB, email, password)
	if err != nil {
		t.Errorf("Expected successful login, but got error: %v", err)
	}

	if tokenExpirationTime.IsZero() {
		t.Error("Expected a valid token expiration time, but got zero value")
	}

	if token == "" {
		t.Error("Expected a valid token, but got an empty string")
	}

	testutils.CleanUpDatabase()
}

func TestLoginUser_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create GORM DB from mock: %v", err)
	}

	email := "error@example.com"
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs(email, 1).
		WillReturnError(errors.New("db query failed"))

	_, _, err = LoginUser(gormDB, email, "password")
	if err == nil || err.Error() != "db query failed" {
		t.Errorf("Expected 'db query failed' error, but got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}

	testutils.CleanUpDatabase()
}

func TestLoginUser_InvalidEmail(t *testing.T) {
	email := "random@y.com"
	password := "password"
	testutils.CreateUser(email, password)

	_, _, err := LoginUser(testutils.TestDB, "invalidemail@y.com", "password")
	if err == nil || err.Error() != "invalid email or password" {
		t.Errorf("Expected 'invalid email or password' error, but got: %v", err)
	}

	testutils.CleanUpDatabase()
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	email := "random@y.com"
	password := "password"
	testutils.CreateUser(email, password)

	_, _, err := LoginUser(testutils.TestDB, email, "wrongpassword")
	if err == nil || err.Error() != "invalid email or password" {
		t.Error("Expected error for incorrect password, but got none")
		}

	testutils.CleanUpDatabase()
}
