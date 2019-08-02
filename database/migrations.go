package database

import (
	"database/sql"
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/gobuffalo/packr/v2"
	"github.com/rubenv/sql-migrate"
)

var migrationBox *packr.Box

// Migrate uses the migrations files to migrate database versions.
func Migrate(db *sql.DB) error {
	utils.Logger.Info("Checking for database migrations...", nil)
	migrationBox = packr.New("migrations", "../migrations")
	migrations := &migrate.PackrMigrationSource{
		Box: migrationBox,
	}

	migrate.SetTable("migration")
	appliedMigrationsCount, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		return err
	}

	if appliedMigrationsCount > 0 {
		utils.Logger.Info(fmt.Sprintf("Applied %d migration(s).", appliedMigrationsCount), nil)
		return nil
	}

	utils.Logger.Info("Database is already up-to-date.", nil)

	return nil
}
