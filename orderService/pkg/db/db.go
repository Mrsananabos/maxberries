package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Schema   string
}

func Connect(cnf ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
		cnf.Schema,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
