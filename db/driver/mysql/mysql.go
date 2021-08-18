package mysql_driver

import (
	"fmt"
	"go-resepee-api/db/repository"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&repository.User{},
		&repository.Material{},
		&repository.RecipeCategory{},
		&repository.Recipe{},
		&repository.RecipeMaterial{},
		&repository.CookStep{},
		&repository.Review{},
		&repository.File{},
	)

	return db
}
