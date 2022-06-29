package dto

type CommunityRequest struct {
	ID                          uint   `json:"ID"`
	Name                        string `json:"name" validate:"required"`
	CommunityImageUrl           string `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string `json:"communityBackgroundImageUrl"`
	Description                 string `json:"description"  validate:"required"`
}
