package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var PostgresDSN string

func init() {
	PostgresDSN = fmt.Sprintf("postgresql://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv("USER_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("HOST_DB"), os.Getenv("PORT_DB"))
	var err error
	DB, err = gorm.Open(postgres.Open(PostgresDSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
