package http

import (
	"fmt"
	"macaiki/internal/domain"
	"macaiki/internal/thread/dto"
	_middL "macaiki/pkg/middleware"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ThreadHandler struct {
	router *echo.Echo
	tu     domain.ThreadUseCase
}

func (th *ThreadHandler) GetThreads(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	trending := c.QueryParam("trending")
	community := c.QueryParam("community")
	forYou := c.QueryParam("forYou")
	keyword := c.QueryParam("keyword")

	var res interface{}
	var err error

	if trending == "true" {
		res, err = th.tu.GetTrendingThreads(uint(userID))
	} else if community == "true" {
		res, err = th.tu.GetThreadsFromFollowedCommunity(uint(userID))
	} else if forYou == "true" {
		res, err = th.tu.GetThreadsFromFollowedUsers(uint(userID))
	} else {
		res, err = th.tu.GetThreads(keyword, uint(userID))
	}

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
	userID, _ := _middL.ExtractTokenUser(c)

	thread := new(dto.ThreadRequest)
	if err := c.Bind(thread); err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	res, err := th.tu.CreateThread(*thread, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (th *ThreadHandler) SetThreadImage(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

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

	err = th.tu.SetThreadImage(img, threadIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) DeleteThread(c echo.Context) error {
	// TODO: Allow admin to delete a thread
	userID, _ := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	if err := th.tu.DeleteThread(threadIDUint, uint(userID)); err != nil {
		return response.ErrorResponse(c, err)

	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) UpdateThread(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

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

	res, err := th.tu.UpdateThread(*thread, threadIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, res)
}

func (th *ThreadHandler) LikeThread(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	err = th.tu.LikeThread(threadIDUint, 1)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) AddThreadComment(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	userID, _ := _middL.ExtractTokenUser(c)
	comment := new(dto.CommentRequest)
	if err := c.Bind(comment); err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	comment.ThreadID = threadIDUint
	comment.UserID = uint(userID)

	err = th.tu.AddThreadComment(*comment)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) GetCommentsByThreadID(c echo.Context) error {
	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	comments, err := th.tu.GetCommentsByThreadID(threadIDUint)

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, comments)
}

func CreateNewThreadHandler(e *echo.Echo, tu domain.ThreadUseCase, JWTSecret string) *ThreadHandler {
	threadHandler := &ThreadHandler{router: e, tu: tu}
	threadHandler.router.POST("/api/v1/threads", threadHandler.CreateThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID", threadHandler.DeleteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.GET("/api/v1/threads", threadHandler.GetThreads, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.GET("/api/v1/threads/:threadID", threadHandler.GetThreadByID)
	threadHandler.router.PUT("/api/v1/threads/:threadID", threadHandler.UpdateThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.PUT("/api/v1/threads/:threadID/images", threadHandler.SetThreadImage, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/likes", threadHandler.LikeThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.PUT("/api/v1/threads/:threadID", threadHandler.UpdateThread)
	threadHandler.router.POST("/api/v1/threads/:threadID/comments", threadHandler.AddThreadComment, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.PUT("/api/v1/threads/:threadID/images", threadHandler.SetThreadImage)
	threadHandler.router.GET("/api/v1/threads/:threadID/comments", threadHandler.GetCommentsByThreadID)
	return threadHandler
}
