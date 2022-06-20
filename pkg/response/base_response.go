package response

import (
	"macaiki/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type baseResponse struct {
	Meta meta
	Data interface{}
}

func SuccessResponse(c echo.Context, data interface{}) error {
	resp := baseResponse{}
	resp.Meta.Code = http.StatusOK
	resp.Meta.Message = "OK"
	resp.Data = data

	return c.JSON(resp.Meta.Code, resp)
}

func ErrorResponse(c echo.Context, err error) error {
	resp := baseResponse{}
	resp.Meta.Code = domain.GetStatusCode(err)
	resp.Meta.Message = err.Error()
	resp.Data = nil

	return c.JSON(resp.Meta.Code, resp)
}
