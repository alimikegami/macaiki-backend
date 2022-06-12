package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	_config "macaiki/config"
	_driver "macaiki/internal/driver"
	_reportCategoryHttpDeliver "macaiki/internal/report_category/delivery/http"
	_reportCategoryRepo "macaiki/internal/report_category/repository/mysql"
	_reportCategoryUsecase "macaiki/internal/report_category/usecase"
	_threadHttpDelivery "macaiki/internal/thread/delivery/http"
	_threadRepo "macaiki/internal/thread/repository/mysql"
	_threadUsecase "macaiki/internal/thread/usecase"
	_userHttpDelivery "macaiki/internal/user/delivery/http"
	_userRepo "macaiki/internal/user/repository/mysql"
	_userUsecase "macaiki/internal/user/usecase"
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
	v := validator.New()

	// setup User
	userRepo := _userRepo.NewMysqlUserRepository(_driver.DB)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, v)
	JWTSecret, err := _config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal("err", err)
	}
	_userHttpDelivery.NewUserHandler(e, userUsecase, JWTSecret.Secret)

	// setup Thread
	threadRepo := _threadRepo.CreateNewThreadRepository(_driver.DB)
	threadUseCase := _threadUsecase.CreateNewThreadUseCase(threadRepo)
	_ = _threadHttpDelivery.CreateNewThreadHandler(e, threadUseCase)

	// setup Report Category
	reportCategoryRepo := _reportCategoryRepo.NewReportCategoryRepository(_driver.DB)
	reportCategoryUsecase := _reportCategoryUsecase.NewReportCategoryUsecase(reportCategoryRepo, v)
	_reportCategoryHttpDeliver.NewReportCategoryHandler(e, reportCategoryUsecase)

	log.Fatal(e.Start(":" + config.ServerPort))
}
