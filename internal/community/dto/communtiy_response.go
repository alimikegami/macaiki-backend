package dto

type CommunityResponse struct {
	ID                          uint   `json:"ID"`
	Name                        string `json:"name"`
	CommunityImageUrl           string `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string `json:"communityBackgroundImageUrl"`
	Description                 string `json:"description"`
}
