package dto

import (
	"time"
)

type UserResponse struct {
	ID              uint   `json:"ID"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profileImageURL"`
	Profession      string `json:"profession"`
	IsFollowed      bool   `json:"isFollowed"`
}

type UserDetailResponse struct {
	ID                 uint      `json:"ID"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	Name               string    `json:"name"`
	ImageUrl           string    `json:"imageUrl"`
	ProfileImageUrl    string    `json:"profileImageURL"`
	BackgroundImageUrl string    `json:"backgroundImageURL"`
	Bio                string    `json:"bio"`
	Profession         string    `json:"profession"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	TotalFollower      int       `json:"totalFollower"`
	TotalFollowing     int       `json:"totalFollowing"`
	TotalPost          int       `json:"totalPost"`
	IsFollowed         bool      `json:"isFollowed"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
