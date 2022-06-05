package request

import "macaiki/domain"

type UserUpdateRequest struct {
	Name      string `json:"name"      validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string
	Role      string `json:"role"     validate:"required"`
	Is_banned int    `json:"isBanned" validate:"required"`
}

func ToUserUpdateRequest(user domain.User) UserUpdateRequest {
	return UserUpdateRequest{
		Name:      user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		Is_banned: user.IsBanned,
	}
}
