package dto

type CommunityModeratorRequest struct {
	UserID      uint `json:"userID"`
	CommunityID uint `json:"communityID"`
}
