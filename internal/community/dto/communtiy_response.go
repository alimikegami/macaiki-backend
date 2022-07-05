package dto

import "macaiki/internal/user/dto"

type CommunityDetailResponse struct {
	ID                          uint   `json:"ID"`
	Name                        string `json:"name"`
	CommunityImageUrl           string `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string `json:"communityBackgroundImageUrl"`
	Description                 string `json:"description"`
	IsFollowed                  int    `json:"isFollowed"`
	IsModerator                 int    `json:"isModerator"`
}

type CommunityUpdateResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommunityAboutResponse struct {
	ID                          uint               `json:"ID"`
	Name                        string             `json:"name"`
	CommunityImageUrl           string             `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string             `json:"communityBackgroundImageUrl"`
	Description                 string             `json:"description"`
	IsFollowed                  int                `json:"isFollowed"`
	IsModerator                 int                `json:"isModerator"`
	TotalModerator              int                `json:"totalModerators"`
	TotalFollower               int                `json:"totalFollowers"`
	Moderator                   []dto.UserResponse `json:"moderators"`
}
