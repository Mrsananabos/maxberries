package migration

import (
	"catalogService/storage/postgres/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func DoMigration() error {
	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	if err = goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("Ошибка при выполнении миграций: %v", err)
	}

	return nil
}
