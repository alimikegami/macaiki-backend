package response

import (
	"macaiki/domain"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"ID"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsBanned  int       `json:"isBanned"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Token struct {
	Token string `json:"token"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListUserResponse(users []domain.User) []UserResponse {
	usersResponse := []UserResponse{}

	for _, val := range users {
		usersResponse = append(usersResponse, ToUserResponse(val))
	}

	return usersResponse
}

func ToTokenResponse(token string) Token {
	return Token{
		token,
	}
}
