package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"macaiki/domain"
	_userHttpDelivery "macaiki/user/delivery/http"
	_userRepo "macaiki/user/repository/mysql"
	_userUsecase "macaiki/user/usecase"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service Run on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// middL
	// e.Use(middL.CORS)
	userRepo := _userRepo.NewMysqlUserRepository(db)

	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
