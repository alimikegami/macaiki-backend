package community

import (
	"macaiki/internal/community/dto"
	dtoCommunity "macaiki/internal/community/dto"
	dtoThread "macaiki/internal/thread/dto"
	"mime/multipart"
)

type CommunityUsecase interface {
	GetAllCommunities(userID int, search string) ([]dtoCommunity.CommunityDetailResponse, error)
	GetCommunity(userID, communityID uint) (dtoCommunity.CommunityDetailResponse, error)
	GetCommunityAbout(userID, communityID uint) (dtoCommunity.CommunityAboutResponse, error)
	StoreCommunity(community dtoCommunity.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dtoCommunity.CommunityRequest, role string) (dtoCommunity.CommunityUpdateResponse, error)
	DeleteCommunity(id uint, role string) error

	FollowCommunity(userID, communityID uint) error
	UnfollowCommunity(userID, communityID uint) error
	SetImage(id uint, img *multipart.FileHeader, role string) (string, error)
	SetBackgroundImage(id uint, img *multipart.FileHeader, role string) (string, error)

	GetThreadCommunity(userID, communityID uint) ([]dtoThread.DetailedThreadResponse, error)
	AddModerator(moderatorReq dtoCommunity.CommunityModeratorRequest, role string) error
	RemoveModerator(moderatorReq dtoCommunity.CommunityModeratorRequest, role string) error

	ReportCommunity(userID, communityID, reportCategoryID uint) error
	GetReports(userID, communityID uint) ([]dto.BriefReportResponse, error)
}
