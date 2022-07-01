package main

import (
	"log"
	"net/http"

	_config "macaiki/config"
	_communityHttpDelivery "macaiki/internal/community/delivery/http"
	_communityRepo "macaiki/internal/community/repository/mysql"
	_communityUsecase "macaiki/internal/community/usecase"
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
	_cloudstorage "macaiki/pkg/cloud_storage"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

	s3Instance := _cloudstorage.CreateNewS3Instance(config.AWSAccessKeyId, config.AWSSecretKey, config.AWSRegion, config.BucketName)
	// setup Repo
	userRepo := _userRepo.NewMysqlUserRepository(_driver.DB)
	reportCategoryRepo := _reportCategoryRepo.NewReportCategoryRepository(_driver.DB)
	threadRepo := _threadRepo.CreateNewThreadRepository(_driver.DB)
	communityRepo := _communityRepo.NewCommunityRepository(_driver.DB)

	// setup usecase
	userUsecase := _userUsecase.NewUserUsecase(userRepo, reportCategoryRepo, v, s3Instance)
	reportCategoryUsecase := _reportCategoryUsecase.NewReportCategoryUsecase(reportCategoryRepo, v)
	threadUseCase := _threadUsecase.CreateNewThreadUseCase(threadRepo, s3Instance)
	communityUsecase := _communityUsecase.NewCommunityUsecase(communityRepo, userRepo, v, s3Instance)

	// setup middleware
	JWTSecret, err := _config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal("err", err)
	}

	// setup route
	_userHttpDelivery.NewUserHandler(e, userUsecase, JWTSecret.Secret)
	_ = _threadHttpDelivery.CreateNewThreadHandler(e, threadUseCase, JWTSecret.Secret)
	_reportCategoryHttpDeliver.NewReportCategoryHandler(e, reportCategoryUsecase, JWTSecret.Secret)
	_communityHttpDelivery.NewCommunityHandler(e, communityUsecase, JWTSecret.Secret)

	log.Fatal(e.Start(":" + config.ServerPort))
}
