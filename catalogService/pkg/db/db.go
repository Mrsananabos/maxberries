package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfingDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Connect(cnf ConfingDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		//"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	//db, err := sqlx.Connect("postgres", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
