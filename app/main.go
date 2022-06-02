package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	_config "macaiki/config"
	_driver "macaiki/driver"
	_userHttpDelivery "macaiki/user/delivery/http"
	_userRepo "macaiki/user/repository/mysql"
	_userUsecase "macaiki/user/usecase"
)

func main() {
	config, err := _config.LoadConfig(".")
	if err != nil {
		log.Fatal("err", err)
	}
	_driver.ConnectDB(
		config.DBConn,
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPass,
		config.DBName,
	)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	userRepo := _userRepo.NewMysqlUserRepository(_driver.DB)
	v := validator.New()
	userUsecase := _userUsecase.NewUserUsecase(userRepo, v)

	JWTSecret, err := _config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal("err", err)
	}

	_userHttpDelivery.NewUserHandler(e, userUsecase, JWTSecret.Secret)

	log.Fatal(e.Start(":" + config.ServerPort))
}
