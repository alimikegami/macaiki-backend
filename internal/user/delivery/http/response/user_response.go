package response

import (
	"macaiki/internal/domain"
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

type UserUpdate struct {
	Name      string `json:"name"      validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string
	Role      string `json:"role"     validate:"required"`
	Is_banned int    `json:"isBanned" validate:"required"`
}

func ToUserUpdate(user domain.User) UserUpdate {
	return UserUpdate{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		Is_banned: user.IsBanned,
	}
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
