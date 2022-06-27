package dto

type ReportCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
