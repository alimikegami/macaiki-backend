package helper

import (
	"macaiki/internal/user/dto"
	"macaiki/internal/user/entity"
)

// Response
func DomainUserToUserResponse(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Name:            user.Name,
		ProfileImageUrl: user.ProfileImageUrl,
		Profession:      user.Profession,
		IsFollowed:      user.IsFollowed,
	}
}

func DomainUserToUserDetailResponse(user entity.User, totalFollowing, totalFollower, totalPost int) dto.UserDetailResponse {
	return dto.UserDetailResponse{
		ID:                 user.ID,
		Username:           user.Username,
		Name:               user.Name,
		ProfileImageUrl:    user.ProfileImageUrl,
		BackgroundImageUrl: user.BackgroundImageUrl,
		Bio:                user.Bio,
		Profession:         user.Profession,
		TotalFollower:      totalFollower,
		TotalFollowing:     totalFollowing,
		TotalPost:          totalPost,
		IsFollowed:         user.IsFollowed,
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
