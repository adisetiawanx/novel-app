package response

import (
	"time"
)

type NovelCreateResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type NovelsGetResponseWrapper struct {
	Novels   []NovelsGetResponse `json:"novels"`
	Total    int64               `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

type ChapterResponse struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type NovelsGetResponse struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Status      string            `json:"status"`
	Country     string            `json:"country"`
	RatingTotal float32           `json:"rating_total"`
	CreatedAt   time.Time         `json:"created_at"`
	Chapters    []ChapterResponse `json:"chapters"`
}
