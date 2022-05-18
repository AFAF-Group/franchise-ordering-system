package config

import (
	"fmt"

	"afaf-group.com/domain/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection(d *Database) (*gorm.DB, error) {
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DatabaseName,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	DatabaseMigration(db)

	return db, nil
}

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Customer{})
	db.AutoMigrate(models.Food{})
}
