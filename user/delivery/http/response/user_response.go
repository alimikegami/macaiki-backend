package response

import (
	"macaiki/domain"
	"time"
)

type UserResponse struct {
	ID        uint                `json:"ID"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	IsBanned  int                 `json:"isBanned"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	Followers []FollowersResponse `json:"followers"`
}

type FollowersResponse struct {
	ID    uint   `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Token struct {
	Token string `json:"token"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Followers: ToListFollowerResponse(user.Followers),
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
