package db

import (
	"authService/configs"
	"authService/internal/permission/model"
	role "authService/internal/role/model"
	user "authService/internal/user/model"
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
}

func Connect(cnf configs.Database) (*gorm.DB, error) {
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

	db.AutoMigrate(&model.Permission{})
	db.AutoMigrate(&role.Role{})
	db.AutoMigrate(&user.User{})

	return db, err
}
