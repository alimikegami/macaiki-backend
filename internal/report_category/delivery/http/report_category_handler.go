package http

import (
	"macaiki/internal/domain"
	"macaiki/pkg/response"

	"github.com/labstack/echo/v4"
)

type ReportCategoryHandler struct {
	rcUsecase domain.ReportCategoryUsecase
}

func NewReportCategoryHandler(e *echo.Echo, rcUsecase domain.ReportCategoryUsecase) {
	rcHandler := ReportCategoryHandler{rcUsecase}

	e.GET("/api/v1/report_categories", rcHandler.GetAllReportCategories)
}

func (rcHandler *ReportCategoryHandler) GetAllReportCategories(c echo.Context) error {

	dtoResponse, err := rcHandler.rcUsecase.GetAllReportCategory()
	if err != nil {
		return response.ErrorResponse(c, domain.ErrInternalServerError)
	}

	return response.SuccessResponse(c, dtoResponse)
}
