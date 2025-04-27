package helper

import (
	"github.com/adisetiawanx/novel-app/internal/dto/response"
	"github.com/adisetiawanx/novel-app/internal/model"
)

func ToNovelsGetResponse(novels []*model.Novel) []response.NovelsGetResponse {
	var responses []response.NovelsGetResponse
	for _, novel := range novels {
		chapters := make([]response.ChapterResponse, 0)
		for _, chapter := range novel.Chapters {
			chapters = append(chapters, response.ChapterResponse{
				Title: chapter.Title,
				Slug:  chapter.Slug,
			})
		}

		responses = append(responses, response.NovelsGetResponse{
			ID:          novel.ID.String(),
			Title:       novel.Title,
			Slug:        novel.Slug,
			Status:      novel.Status,
			Country:     novel.Country,
			RatingTotal: novel.RatingTotal,
			CreatedAt:   novel.CreatedAt,
			Chapters:    chapters,
		})
	}
	return responses
}

func ToMediaUploadNovelCoverResponse(media []*model.Media) []response.MediaUploadNovelCoverResponse {
	var result []response.MediaUploadNovelCoverResponse
	for _, m := range media {
		result = append(result, response.MediaUploadNovelCoverResponse{
			ID:        m.ID.String(),
			Url:       m.Url,
			Name:      m.Name,
			Size:      m.Size,
			Type:      m.Type,
			CreatedAt: m.CreatedAt,
		})
	}
	return result
}
