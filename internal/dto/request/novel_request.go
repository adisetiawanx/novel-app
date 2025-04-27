package request

import "github.com/adisetiawanx/novel-app/internal/model"

type NovelCreateRequest struct {
	PostStatus       string              `json:"post_status" validate:"required,oneof=publish draft"`
	Title            string              `json:"title" validate:"required"`
	Slug             string              `json:"slug" validate:"required"`
	AlternativeTitle string              `json:"alternative_title"`
	Synopsis         string              `json:"synopsis"`
	Status           string              `json:"status" validate:"required,oneof=ongoing completed hiatus"`
	ReleaseYear      int16               `json:"release_year" validate:"required"`
	Country          string              `json:"country" validate:"required,oneof=china korea japan indonesia"`
	Genres           []*model.Genre      `json:"genres" `
	Authors          []*model.Author     `json:"authors"`
	Translators      []*model.Translator `json:"translators"`
}
