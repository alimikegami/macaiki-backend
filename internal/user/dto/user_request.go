package dto

type UserRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}

type UserUpdateRequest struct {
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	Profession string `json:"profession"`
}

type UserChangePasswordRequest struct {
	NewPassword          string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserReportRequest struct {
	ReportCategoryID uint `json:"reportCategoryID"`
}

type SendOTPRequest struct {
	Email string `json:"email"`
	Link  string `json:"link"`
}

type EmailMessage struct {
	EmailDest    string `json:"email_destination"`
	EmailBody    string `json:"email_body"`
	EmailSubject string `json:"email_subject"`
}
