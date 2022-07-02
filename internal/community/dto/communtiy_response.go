package dto

type CommunityDetailResponse struct {
	ID                          uint   `json:"ID"`
	Name                        string `json:"name"`
	CommunityImageUrl           string `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string `json:"communityBackgroundImageUrl"`
	Description                 string `json:"description"`
	IsFollowed                  bool   `json:"isFollowed"`
}

type CommunityResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
