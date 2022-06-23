package dto

type ThreadRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	CommunityID uint   `json:"communityID"`
}
