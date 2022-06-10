package driver

import (
	"fmt"
	"log"
	"macaiki/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(driver, host, port, username, password, name string) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	if driver == "MYSQL" {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("err", err)
		}
		InitialMigration(DB)
	}
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(
		&domain.User{},
		&domain.Thread{},
		&domain.ReportCategory{},
	)
}
