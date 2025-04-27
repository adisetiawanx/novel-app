package response

import "time"

type MediaUploadNovelCoverResponseWrapper struct {
	Media []MediaUploadNovelCoverResponse `json:"media"`
	Total int                             `json:"total"`
}
type MediaUploadNovelCoverResponse struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
