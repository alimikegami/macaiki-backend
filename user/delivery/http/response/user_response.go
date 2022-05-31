package response

import "macaiki/domain"

type UserResponse struct {
	ID        uint
	Name      string
	Email     string
	Is_banned int
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Is_banned: user.Is_banned,
	}
}

func ToListUserResponse(users []domain.User) []UserResponse {
	usersResponse := []UserResponse{}

	for _, val := range users {
		usersResponse = append(usersResponse, ToUserResponse(val))
	}

	return usersResponse
}
