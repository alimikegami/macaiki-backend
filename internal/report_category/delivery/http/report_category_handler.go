package http

import (
	"macaiki/internal/domain"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReportCategoryHandler struct {
	rcUsecase domain.ReportCategoryUsecase
}

func NewReportCategoryHandler(e *echo.Echo, rcUsecase domain.ReportCategoryUsecase) {
	rcHandler := ReportCategoryHandler{rcUsecase}

	e.GET("/api/v1/report-categories", rcHandler.GetAllReportCategories)
	e.GET("/api/v1/report-categories/:reportCategoryID", rcHandler.GetReportCategory)
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
