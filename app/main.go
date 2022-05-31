package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_driver "macaiki/driver"
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

	db := _driver.ConnectDB(dbHost, dbPort, dbUser, dbPass, dbName)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// middL
	// e.Use(middL.CORS)
	userRepo := _userRepo.NewMysqlUserRepository(db)
	validator := validator.New()
	userUsecase := _userUsecase.NewUserUsecase(userRepo, validator)
	_userHttpDelivery.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
