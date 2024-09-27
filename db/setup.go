package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
    databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
				return nil, fmt.Errorf("DATABASE_URL must be set")
		}

		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
				return nil, fmt.Errorf("error opening database: %v", err)
		}

		err = db.Ping()
		if err != nil {
				return nil, fmt.Errorf("error pinging database: %v", err)
		}

		if err := migrationsApplied(db); err != nil {
			return nil, fmt.Errorf("unapplied migrations present: %v", err)
		}

		fmt.Println("Successfully connected to the database!")
		return db, nil
}

func migrationsApplied(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
			return fmt.Errorf("could not start the migration driver: %v", err)
	}

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
			return fmt.Errorf("MIGRATION_PATH must be set")
	}

	m, err := migrate.NewWithDatabaseInstance(
			migrationPath,
			"postgres", driver,
	)
	if err != nil {
			return fmt.Errorf("could not initialize migration: %v", err)
	}

	version, dirty, err := m.Version()
	if err == migrate.ErrNilVersion {
			return fmt.Errorf("no migrations applied")
	}
	if err != nil {
			return fmt.Errorf("error checking migration version: %v", err)
	}

	if dirty {
			return fmt.Errorf("the database is in a dirty state, migration version: %d", version)
	}

	latestMigrationVersion, err := getLatestMigrationVersion(migrationPath)
	if err != nil {
			return fmt.Errorf("could not get the latest migration version: %v", err)
	}

	if version < latestMigrationVersion {
			return fmt.Errorf("not all migrations have been applied, current version: %d, latest version: %d", version, latestMigrationVersion)
	}

	return nil
}

func getLatestMigrationVersion(migrationsPath string) (uint, error) {
	var latestVersion uint = 0

	migrationsPath = strings.TrimPrefix(migrationsPath, "file://")
  
	err := filepath.Walk(migrationsPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			if strings.HasSuffix(info.Name(), ".up.sql") {
					parts := strings.Split(info.Name(), "_")
					if len(parts) > 0 {
							versionStr := parts[0]
							version, err := strconv.ParseUint(versionStr, 10, 64)
							if err != nil {
									return nil
							}
							if uint(version) > latestVersion {
									latestVersion = uint(version)
							}
					}
			}
			return nil
	})

	if err != nil {
			return 0, err
	}

	return latestVersion, nil
}