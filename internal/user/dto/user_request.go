package dto

type UserRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
	// Name               string `json:"name"`
	// ProfileImageUrl    string `json:"profileImageURL"`
	// BackgroundImageUrl string `json:"backgroundImageURL"`
	// Bio                string `json:"bio"`
	// Proffesion         string `json:"proffesion"`
	// Role               string `json:"role"`
	// IsBanned           bool   `json:"isBanned"`
}

type UpdateUserRequest struct {
	Email              string `json:"email"`
	Username           string `json:"username"`
	Name               string `json:"name"`
	ProfileImageUrl    string `json:"profileImageURL"`
	BackgroundImageUrl string `json:"backgroundImageURL"`
	Bio                string `json:"bio"`
	Proffesion         string `json:"proffesion"`
	Role               string `json:"role"`
	IsBanned           bool   `json:"isBanned"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}
