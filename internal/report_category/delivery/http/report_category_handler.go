package http

import (
	"macaiki/internal/domain"
	"macaiki/internal/report_category/dto"
	_middL "macaiki/pkg/middleware"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ReportCategoryHandler struct {
	rcUsecase domain.ReportCategoryUsecase
	JWTSecret string
}

func NewReportCategoryHandler(e *echo.Echo, rcUsecase domain.ReportCategoryUsecase, JWTSecret string) {
	rcHandler := ReportCategoryHandler{rcUsecase, JWTSecret}
	e.POST("api/v1/report-categories", rcHandler.CreateReportCategory, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/report-categories", rcHandler.GetAllReportCategories)
	e.GET("/api/v1/report-categories/:reportCategoryID", rcHandler.GetReportCategory)
	e.PUT("/api/v1/report-categories/:reportCategoryID", rcHandler.UpdateReportCategory, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/report-categories/:reportCategoryID", rcHandler.DeleteReportCategory, middleware.JWT([]byte(JWTSecret)))
}

func (rcHandler *ReportCategoryHandler) CreateReportCategory(c echo.Context) error {
	rcReq := dto.ReportCategoryRequest{}
	c.Bind(&rcReq)

	_, role := _middL.ExtractTokenUser(c)
	err := rcHandler.rcUsecase.CreateReportCategory(rcReq, role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (rcHandler *ReportCategoryHandler) GetAllReportCategories(c echo.Context) error {
	dtoResponse, err := rcHandler.rcUsecase.GetAllReportCategory()
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, dtoResponse)
}

func (rcHandler *ReportCategoryHandler) GetReportCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("reportCategoryID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	dtoResponse, err := rcHandler.rcUsecase.GetReportCategory(uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, dtoResponse)
}

func (rcHandler *ReportCategoryHandler) UpdateReportCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("reportCategoryID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	rcReq := dto.ReportCategoryRequest{}
	c.Bind(&rcReq)

	_, role := _middL.ExtractTokenUser(c)
	err = rcHandler.rcUsecase.UpdateReportCategory(rcReq, uint(id), role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (rcHandler *ReportCategoryHandler) DeleteReportCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("reportCategoryID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	_, role := _middL.ExtractTokenUser(c)
	err = rcHandler.rcUsecase.DeleteReportCategory(uint(id), role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}
