package dto

type UserRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}

type UpdateUserRequest struct {
	Username           string `json:"username"`
	Name               string `json:"name"`
	ProfileImageUrl    string `json:"profileImageURL"`
	BackgroundImageUrl string `json:"backgroundImageURL"`
	Bio                string `json:"bio"`
	Profession         string `json:"profession"`
	Role               string `json:"role"`
	IsBanned           bool   `json:"isBanned"`
}

type ChangePasswordUserRequest struct {
	NewPassword          string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserReportRequest struct {
	UserID           uint `json:"userID"`
	ReportCategoryID uint `json:"reportCategoryID"`
}
