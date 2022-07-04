package dto

type UserResponse struct {
	ID              uint   `json:"ID"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profileImageURL"`
	Profession      string `json:"profession"`
	IsFollowed      int    `json:"isFollowed"`
}

type UserDetailResponse struct {
	ID                 uint   `json:"ID"`
	Username           string `json:"username"`
	Name               string `json:"name"`
	ProfileImageUrl    string `json:"profileImageURL"`
	BackgroundImageUrl string `json:"backgroundImageURL"`
	Bio                string `json:"bio"`
	Profession         string `json:"profession"`
	TotalFollower      int    `json:"totalFollower"`
	TotalFollowing     int    `json:"totalFollowing"`
	TotalPost          int    `json:"totalPost"`
	IsFollowed         int    `json:"isFollowed"`
	IsMine             int    `json:"isMine"`
}

type UserUpdateResponse struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	Profession string `json:"profession"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
