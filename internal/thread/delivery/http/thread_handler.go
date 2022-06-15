package http

import (
	"fmt"
	"macaiki/internal/domain"
	"macaiki/internal/thread/dto"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ThreadHandler struct {
	router *echo.Echo
	tu     domain.ThreadUseCase
}

func (th *ThreadHandler) GetThreads(c echo.Context) error {
	res, err := th.tu.GetThreads()
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (th *ThreadHandler) GetThreadByID(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	res, err := th.tu.GetThreadByID(threadIDUint)
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
		return response.ErrorResponse(c, err)
	}

	res, err := th.tu.CreateThread(*thread, 1)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (th *ThreadHandler) SetThreadImage(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	img, err := c.FormFile("threadImg")
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	err = th.tu.SetThreadImage(img, threadIDUint)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) DeleteThread(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	if err := th.tu.DeleteThread(threadIDUint); err != nil {
		return response.ErrorResponse(c, err)

	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) UpdateThread(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	thread := new(dto.ThreadRequest)
	if err := c.Bind(thread); err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	res, err := th.tu.UpdateThread(*thread, threadIDUint, 1)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, res)
}

func CreateNewThreadHandler(e *echo.Echo, tu domain.ThreadUseCase) *ThreadHandler {
	threadHandler := &ThreadHandler{router: e, tu: tu}
	threadHandler.router.POST("/api/v1/threads", threadHandler.CreateThread)
	threadHandler.router.DELETE("/api/v1/threads/:threadID", threadHandler.DeleteThread)
	threadHandler.router.GET("/api/v1/threads", threadHandler.GetThreads)
	threadHandler.router.GET("/api/v1/threads/:threadID", threadHandler.GetThreadByID)
	threadHandler.router.PUT("/api/v1/threads/:threadID", threadHandler.UpdateThread)
	threadHandler.router.PUT("/api/v1/threads/:threadID/images", threadHandler.SetThreadImage)
	return threadHandler
}
