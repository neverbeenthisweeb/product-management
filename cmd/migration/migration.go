package main

import (
	"log"
	"productmanagement/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Config
	cfg := config.NewConfigImpl()
	sourceURL := cfg.GetString("SOURCE_URL")
	dbURL := cfg.GetString("DATABASE_URL")

	// Migration
	m, err := migrate.New(
		"file://"+sourceURL,
		"mysql://"+dbURL,
	)
	if err != nil {
		panic("Failed to initiate migration: " + err.Error())
	}

	err = m.Down()
	if err != nil {
		log.Println("Failed to migrate down: " + err.Error() + ". About to force version")
		if m.Force(1) != nil {
			panic("Failed to force version: " + err.Error())
		}
	}

	err = m.Up()
	if err != nil {
		panic("Failed to migrate up: " + err.Error())
	}
}
