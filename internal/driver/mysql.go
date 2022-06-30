package driver

import (
	"fmt"
	"log"
	community "macaiki/internal/community/entity"
	reportCategoryEntity "macaiki/internal/report_category/entity"
	threadEntity "macaiki/internal/thread/entity"
	userEntity "macaiki/internal/user/entity"

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
		&reportCategoryEntity.ReportCategory{},
		&community.Community{},
		&userEntity.User{},
		&userEntity.UserReport{},
		&threadEntity.Thread{},
		&threadEntity.ThreadLikes{},
		&threadEntity.ThreadFollower{},
		&userEntity.FollowedCommunity{},
		&threadEntity.Comment{},
	)
}
