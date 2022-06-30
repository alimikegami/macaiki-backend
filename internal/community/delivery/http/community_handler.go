package http

import (
	"macaiki/internal/community/dto"
	"macaiki/internal/domain"
	"macaiki/pkg/response"
	"strconv"

	_middL "macaiki/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CommunityHandler struct {
	communityUsecase domain.CommunityUsecase
	JWTSecret        string
}

func NewCommunityHandler(e *echo.Echo, communityUsecase domain.CommunityUsecase, JWTSecret string) {
	communityHandler := &CommunityHandler{
		communityUsecase: communityUsecase,
		JWTSecret:        JWTSecret,
	}

	e.POST("api/v1/communities", communityHandler.CreateCommunity, middleware.JWT([]byte(JWTSecret)))
	// e.GET("api/v1/communities", communityHandler.GetAllCommunity)
	e.GET("api/v1/communities", communityHandler.GetAllCommunityWithDetail, middleware.JWT([]byte(JWTSecret)))
	e.GET("api/v1/communities/:communityID", communityHandler.GetCommunity)
	e.PUT("api/v1/communities/:communityID", communityHandler.UpdateCommunity, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("api/v1/communities/:communityID", communityHandler.DeleteCommunity, middleware.JWT([]byte(JWTSecret)))

	e.POST("api/v1/curent-user/community-followers/:communityID", communityHandler.FollowCommunity, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("api/v1/curent-user/community-followers/:communityID", communityHandler.UnfollowCommunity, middleware.JWT([]byte(JWTSecret)))
}

func (communityHandler *CommunityHandler) CreateCommunity(c echo.Context) error {
	communityReq := dto.CommunityRequest{}
	c.Bind(&communityReq)

	_, role := _middL.ExtractTokenUser(c)
	err := communityHandler.communityUsecase.StoreCommunity(communityReq, role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

// TODO implement this function when the user is logged in as anonymous
// func (communityHandler *CommunityHandler) GetAllCommunity(c echo.Context) error {
// 	search := c.QueryParam("search")
// 	communitiesResp, err := communityHandler.communityUsecase.GetAllCommunity(search)
// 	if err != nil {
// 		return response.ErrorResponse(c, err)
// 	}

// 	return response.SuccessResponse(c, communitiesResp)
// }

func (communityHandler *CommunityHandler) GetAllCommunityWithDetail(c echo.Context) error {
	search := c.QueryParam("search")

	userID, _ := _middL.ExtractTokenUser(c)

	communitiesResp, err := communityHandler.communityUsecase.GetAllCommunityDetail(userID, search)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, communitiesResp)
}

func (communityHandler *CommunityHandler) GetCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	communityResp, err := communityHandler.communityUsecase.GetCommunity(uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, communityResp)
}

func (communityHandler *CommunityHandler) UpdateCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	var communityReq dto.CommunityRequest
	c.Bind(&communityReq)

	_, role := _middL.ExtractTokenUser(c)
	err = communityHandler.communityUsecase.UpdateCommunity(uint(id), communityReq, role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (communityHandler *CommunityHandler) DeleteCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	_, role := _middL.ExtractTokenUser(c)
	err = communityHandler.communityUsecase.DeleteCommunity(uint(id), role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (communityHandler *CommunityHandler) FollowCommunity(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	communityID, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	err = communityHandler.communityUsecase.FollowCommunity(uint(userID), uint(communityID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (communityHandler *CommunityHandler) UnfollowCommunity(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	communityID, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	err = communityHandler.communityUsecase.UnfollowCommunity(uint(userID), uint(communityID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}
