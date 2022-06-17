package dto

import (
	"time"
)

type UserResponse struct {
	ID                 uint      `json:"ID"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	Name               string    `json:"name"`
	ImageUrl           string    `json:"imageUrl"`
	ProfileImageUrl    string    `json:"profileImageURL"`
	BackgroundImageUrl string    `json:"backgroundImageURL"`
	Bio                string    `json:"bio"`
	Proffesion         string    `json:"proffesion"`
	Role               string    `json:"role"`
	IsBanned           bool      `json:"isBanned"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
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
	Proffesion         string    `json:"proffesion"`
	Role               string    `json:"role"`
	IsBanned           bool      `json:"isBanned"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	TotalFollower      int       `json:"totalFollower"`
	TotalFollowing     int       `json:"totalFollowing"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
