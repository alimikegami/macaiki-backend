package driver

import (
	"fmt"
	"log"
	communityEntity "macaiki/internal/community/entity"
	notifEntity "macaiki/internal/notification/entity"
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
		&communityEntity.Community{},
		&userEntity.User{},
		&userEntity.UserReport{},
		&notifEntity.Notification{},
		&communityEntity.CommunityReport{},
		&threadEntity.Thread{},
		&threadEntity.ThreadUpvote{},
		&threadEntity.ThreadFollower{},
		&threadEntity.Comment{},
		&threadEntity.CommentLikes{},
		&threadEntity.ThreadDownvote{},
		&threadEntity.ThreadReport{},
		&threadEntity.CommentReport{},
	)
}
