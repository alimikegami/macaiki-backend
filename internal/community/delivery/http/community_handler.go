package http

import (
	"macaiki/internal/community"
	"macaiki/internal/community/dto"
	"macaiki/pkg/response"
	"macaiki/pkg/utils"
	"strconv"

	_middL "macaiki/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CommunityHandler struct {
	communityUsecase community.CommunityUsecase
	JWTSecret        string
}

func NewCommunityHandler(e *echo.Echo, communityUsecase community.CommunityUsecase, JWTSecret string) {
	communityHandler := &CommunityHandler{
		communityUsecase: communityUsecase,
		JWTSecret:        JWTSecret,
	}

	e.POST("api/v1/communities", communityHandler.CreateCommunity, middleware.JWT([]byte(JWTSecret)))
	e.GET("api/v1/communities", communityHandler.GetAllCommunity)
	e.GET("api/v1/communities/:communityID", communityHandler.GetCommunity)
	e.PUT("api/v1/communities/:communityID", communityHandler.UpdateCommunity, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("api/v1/communities/:communityID", communityHandler.DeleteCommunity, middleware.JWT([]byte(JWTSecret)))
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

func (communityHandler *CommunityHandler) GetAllCommunity(c echo.Context) error {
	search := c.QueryParam("search")
	communitiesResp, err := communityHandler.communityUsecase.GetAllCommunity(search)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, communitiesResp)
}

func (CommunityHandler *CommunityHandler) GetCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	communityResp, err := CommunityHandler.communityUsecase.GetCommunity(uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, communityResp)
}

func (CommunityHandler *CommunityHandler) UpdateCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	var communityReq dto.CommunityRequest
	c.Bind(&communityReq)

	_, role := _middL.ExtractTokenUser(c)
	err = CommunityHandler.communityUsecase.UpdateCommunity(uint(id), communityReq, role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (CommunityHandler *CommunityHandler) DeleteCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("communityID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	_, role := _middL.ExtractTokenUser(c)
	err = CommunityHandler.communityUsecase.DeleteCommunity(uint(id), role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}
