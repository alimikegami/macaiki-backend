package helper

import (
	"macaiki/internal/domain"
	"macaiki/internal/user/dto"
)

// Response
func DomainUserToUserResponse(user domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:                 user.ID,
		Email:              user.Email,
		Username:           user.Username,
		Name:               user.Name,
		ProfileImageUrl:    user.ProfileImageUrl,
		BackgroundImageUrl: user.BackgroundImageUrl,
		Bio:                user.Bio,
		Profession:         user.Profession,
		Role:               user.Role,
		IsBanned:           user.IsBanned,
		CreatedAt:          user.CreatedAt,
		UpdatedAt:          user.UpdatedAt,
	}
}

func DomainUserToUserDetailResponse(user domain.User, totalFollowing, totalFollower, totalPost int) dto.UserDetailResponse {
	return dto.UserDetailResponse{
		ID:                 user.ID,
		Email:              user.Email,
		Username:           user.Username,
		Name:               user.Name,
		ProfileImageUrl:    user.ProfileImageUrl,
		BackgroundImageUrl: user.BackgroundImageUrl,
		Bio:                user.Bio,
		Profession:         user.Profession,
		Role:               user.Role,
		IsBanned:           user.IsBanned,
		CreatedAt:          user.CreatedAt,
		UpdatedAt:          user.UpdatedAt,
		TotalFollower:      totalFollower,
		TotalFollowing:     totalFollowing,
		TotalPost:          totalPost,
	}
}

func DomainUserToListUserResponse(users []domain.User) []dto.UserResponse {
	usersResponse := []dto.UserResponse{}

	for _, val := range users {
		usersResponse = append(usersResponse, DomainUserToUserResponse(val))
	}

	return usersResponse
}

func ToLoginResponse(token string) dto.LoginResponse {
	return dto.LoginResponse{
		Token: token,
	}
}
