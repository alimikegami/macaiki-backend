package helper

import (
	"macaiki/internal/user/dto"
	"macaiki/internal/user/entity"
)

// Response
func DomainUserToUserResponse(user entity.User) dto.UserResponse {
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

func DomainUserToUserDetailResponse(user entity.User, totalFollowing, totalFollower, totalPost int) dto.UserDetailResponse {
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

func DomainUserToListUserResponse(users []entity.User) []dto.UserResponse {
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
