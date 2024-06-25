package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Grafiters/archive/db/config"
)

func Create() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func Drop() {
	db, err := config.ConfigDBBase()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	error := config.DeleteDatabase(db)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("DONE Delete Database")
}

func Migrate(direction string, table string) error {
	db, err := config.ConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	path := "db/migration"
	if table != "" {
		path = fmt.Sprintf("db/migration/%s", table)
	}

	version, err := getCurrentMigrationVersion()
	if err != nil {
		log.Fatal(err)
	}

	migrateCmd := exec.Command("migrate", "-database", fmt.Sprintf("postgres://%s", config.DnsConnection), "-path", path, direction)
	if path == "" {
		migrateCmd = exec.Command("migrate", "-database", fmt.Sprintf("postgres://%s", config.DnsConnection), "-path", path, direction, fmt.Sprintf("%d", version))
	}
	if direction == "down" {
		migrateCmd.Stdin = strings.NewReader("y\n")
	}
	migrateCmd.Stdout = os.Stdout
	migrateCmd.Stderr = os.Stderr

	if err := migrateCmd.Run(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	return nil
}

func getCurrentMigrationVersion() (int, error) {
	db, err := config.ConfigDB()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	var version int
	query := `SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1`

	// Execute query
	err = db.QueryRow(query).Scan(&version)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No migrations have been applied yet")
			return 0, nil
		}
		return 0, fmt.Errorf("failed to fetch current migration version: %w", err)
	}

	return version, nil
}
