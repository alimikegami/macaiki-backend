package dto

import (
	"macaiki/internal/domain"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"ID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsBanned  int       `json:"isBanned"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserDetailResponse struct {
	ID              uint                `json:"ID"`
	Username        string              `json:"username"`
	Email           string              `json:"email"`
	IsBanned        int                 `json:"isBanned"`
	ImageUrl        string              `json:"imageUrl"`
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
	TotalFollowers  int                 `json:"totalFollowers"`
	TotalFollowings int                 `json:"totalFollowing"`
	Followers       []FollowersResponse `json:"followers"`
	Followings      []FollowersResponse `json:"followings"`
}

type FollowersResponse struct {
	ID       uint   `json:"ID"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

type Token struct {
	Token string `json:"token"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		ImageUrl:  user.ImageUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserDetailResponse(user domain.User, followings []domain.User) UserDetailResponse {
	return UserDetailResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		IsBanned:        user.IsBanned,
		ImageUrl:        user.ImageUrl,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		TotalFollowers:  len(user.Followers),
		TotalFollowings: len(followings),
		Followers:       ToListFollowerResponse(user.Followers),
		Followings:      ToListFollowerResponse(followings),
	}
}

func ToListUserResponse(users []domain.User) []UserResponse {
	usersResponse := []UserResponse{}

	for _, val := range users {
		usersResponse = append(usersResponse, ToUserResponse(val))
	}

	return usersResponse
}

func ToListFollowerResponse(followers []domain.User) []FollowersResponse {
	followersResponse := []FollowersResponse{}

	for _, val := range followers {
		temp := FollowersResponse{
			val.ID,
			val.Username,
			val.Email,
		}
		followersResponse = append(followersResponse, temp)
	}

	return followersResponse
}

func ToTokenResponse(token string) Token {
	return Token{
		token,
	}
}
