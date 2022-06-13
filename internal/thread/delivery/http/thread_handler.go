package http

import (
	"fmt"
	"macaiki/internal/domain"
	"macaiki/internal/thread/dto"
	"macaiki/pkg/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ThreadHandler struct {
	router *echo.Echo
	tu     domain.ThreadUseCase
}

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func (th *ThreadHandler) GetThreads(c echo.Context) error {
	res, err := th.tu.GetThreads()
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (th *ThreadHandler) CreateThread(c echo.Context) error {
	thread := new(dto.ThreadRequest)
	if err := c.Bind(thread); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
	}

	if err := th.tu.CreateThread(*thread, 1); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Data:    nil,
		Message: nil,
	})
}

func (th *ThreadHandler) DeleteThread(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
	}
	threadIDUint := uint(u64)
	if err := th.tu.DeleteThread(threadIDUint); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Data:    nil,
		Message: nil,
	})
}

func CreateNewThreadHandler(e *echo.Echo, tu domain.ThreadUseCase) *ThreadHandler {
	threadHandler := &ThreadHandler{router: e, tu: tu}
	threadHandler.router.POST("/api/v1/threads", threadHandler.CreateThread)
	threadHandler.router.DELETE("/api/v1/threads/:threadID", threadHandler.DeleteThread)
	threadHandler.router.GET("/api/v1/threads", threadHandler.GetThreads)
	return threadHandler
}
