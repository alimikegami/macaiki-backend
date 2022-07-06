package http

import (
	"fmt"
	"macaiki/internal/thread"
	"macaiki/internal/thread/dto"
	_middL "macaiki/pkg/middleware"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ThreadHandler struct {
	router *echo.Echo
	tu     thread.ThreadUseCase
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
	userID, role := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)
	if err := th.tu.DeleteThread(threadIDUint, uint(userID), role); err != nil {
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

func (th *ThreadHandler) UpvoteThread(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	err = th.tu.UpvoteThread(threadIDUint, uint(userID))
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

func (th *ThreadHandler) LikeComment(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	commentID := c.Param("commentID")
	u64, err := strconv.ParseUint(commentID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	commentIDUint := uint(u64)

	err = th.tu.LikeComment(commentIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) UnlikeComment(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	commentID := c.Param("commentID")
	u64, err := strconv.ParseUint(commentID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	commentIDUint := uint(u64)

	err = th.tu.UnlikeComment(commentIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) DownvoteThread(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	err = th.tu.DownvoteThread(threadIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) UndoDownvoteThread(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	err = th.tu.UndoDownvoteThread(threadIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) UndoUpvoteThread(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	err = th.tu.UndoUpvoteThread(threadIDUint, uint(userID))
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) DeleteComment(c echo.Context) error {
	// TODO: Allow admin to delete a thread
	userID, role := _middL.ExtractTokenUser(c)

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	threadIDUint := uint(u64)

	commentID := c.Param("commentID")
	u64, err = strconv.ParseUint(commentID, 10, 32)
	commentIDUint := uint(u64)
	if err := th.tu.DeleteComment(commentIDUint, threadIDUint, uint(userID), role); err != nil {
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) CreateThreadReport(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	threadReport := new(dto.ThreadReportRequest)
	if err := c.Bind(threadReport); err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	threadID := c.Param("threadID")
	u64, err := strconv.ParseUint(threadID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	threadIDUint := uint(u64)

	threadReport.ThreadID = threadIDUint
	threadReport.UserID = uint(userID)

	err = th.tu.CreateThreadReport(*threadReport)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (th *ThreadHandler) CreateCommentReport(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	commentReport := new(dto.CommentReportRequest)
	if err := c.Bind(commentReport); err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	commentID := c.Param("commentID")
	u64, err := strconv.ParseUint(commentID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}
	commentIDUint := uint(u64)

	commentReport.CommentID = commentIDUint
	commentReport.UserID = uint(userID)

	err = th.tu.CreateCommentReport(*commentReport)
	if err != nil {
		fmt.Println(err)
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func CreateNewThreadHandler(e *echo.Echo, tu thread.ThreadUseCase, JWTSecret string) *ThreadHandler {
	threadHandler := &ThreadHandler{router: e, tu: tu}
	threadHandler.router.POST("/api/v1/threads", threadHandler.CreateThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID", threadHandler.DeleteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.GET("/api/v1/threads", threadHandler.GetThreads, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.GET("/api/v1/threads/:threadID", threadHandler.GetThreadByID)
	threadHandler.router.PUT("/api/v1/threads/:threadID", threadHandler.UpdateThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.PUT("/api/v1/threads/:threadID/images", threadHandler.SetThreadImage, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/upvotes", threadHandler.UpvoteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/comments", threadHandler.AddThreadComment, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.GET("/api/v1/threads/:threadID/comments", threadHandler.GetCommentsByThreadID)
	threadHandler.router.POST("/api/v1/threads/:threadID/comments/:commentID/likes", threadHandler.LikeComment, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/downvotes", threadHandler.DownvoteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID/upvotes", threadHandler.UndoUpvoteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID/downvotes", threadHandler.UndoDownvoteThread, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID/comments/:commentID/likes", threadHandler.UnlikeComment, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.DELETE("/api/v1/threads/:threadID/comments/:commentID", threadHandler.DeleteComment, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/reports", threadHandler.CreateThreadReport, middleware.JWT([]byte(JWTSecret)))
	threadHandler.router.POST("/api/v1/threads/:threadID/comments/:commentID/reports", threadHandler.CreateCommentReport, middleware.JWT([]byte(JWTSecret)))

	return threadHandler
}
