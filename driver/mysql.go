package driver

import (
	"fmt"
	"macaiki/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(host, port, username, password, name string) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration(DB)
	return DB
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(
		&domain.User{},
	)
}
