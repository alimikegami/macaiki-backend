package dto

type CommunityRequest struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"  validate:"required"`
}
