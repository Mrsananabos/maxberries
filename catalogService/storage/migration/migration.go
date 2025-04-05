package migration

import (
	"catalogService/storage/postgres/config"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

func DoMigration() error {
	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	if err = goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Ошибка при выполнении миграций: %v", err)
	}

	return nil
}
