package testutils

import (
	"coach/db"
	"coach/models"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func GetProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}

		wd = filepath.Dir(wd)
		if wd == "/" || wd == "." {
			break
		}
	}

	return "", err
}

func SetupTestEnv(m *testing.M) {
	e := godotenv.Load("../.env")
	if e != nil {
			log.Printf("Could not load .env file: %v", e)
	}

	projectRoot, e := GetProjectRoot()
	if e != nil {
		log.Fatalf("Failed to find project root: %v", e)
	}

	migrationPath := "file://" + filepath.Join(projectRoot, "db", "migrations")
	os.Setenv("MIGRATION_PATH", migrationPath)
	var err error
	TestDB, err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the test database: %v", err)
	}

}

func CleanUpDatabase() {
	if err := TestDB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatalf("Failed to clean up database: %v", err)
	}
}

func RunTests(m *testing.M) {
	code := m.Run()

	CleanUpDatabase()

	os.Exit(code)
}

func CreateUser(email, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user := models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := TestDB.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create user in the test database: %v", err)
	}
	
}

func SetupTestDB(t *testing.T, models ...interface{}) *gorm.DB {
	t.Helper()

	gormDB, err := db.InitDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return gormDB
}