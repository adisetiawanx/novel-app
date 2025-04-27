package service

import (
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/google/uuid"
)

type NovelService interface {
	Save(novel *model.Novel) (*model.Novel, error)
	FindAll(page int, pageSize int, title string, status string, country string, genres []string, sortBy string, order string) ([]*model.Novel, int64, error)
}

type novelServiceImpl struct {
	repository.NovelRepository
}

func NewNovelService(novelRepository repository.NovelRepository) NovelService {
	return &novelServiceImpl{
		NovelRepository: novelRepository,
	}
}

func (service *novelServiceImpl) Save(novel *model.Novel) (*model.Novel, error) {
	isSlugExist, err := service.NovelRepository.IsSlugExist(novel.Slug)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	if isSlugExist {
		novel.Slug = helper.GenerateUniqueSlug(novel.Slug)
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return nil, helper.NewInternalServerError()
	}
	novel.ID = newId

	for index, genre := range novel.Genres {
		newId, err := uuid.NewV7()
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		genre.ID = newId
		novel.Genres[index] = genre
	}

	for index, author := range novel.Authors {
		newId, err := uuid.NewV7()
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		author.ID = newId
		novel.Authors[index] = author
	}

	for index, translator := range novel.Translators {
		newId, err := uuid.NewV7()
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		translator.ID = newId
		novel.Translators[index] = translator
	}

	_, err = service.NovelRepository.Save(novel)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return novel, nil
}

func (service *novelServiceImpl) FindAll(page int, pageSize int, title string, status string, country string, genres []string, sortBy string, order string) ([]*model.Novel, int64, error) {
	novels, totalRows, err := service.NovelRepository.FindAll(page, pageSize, title, status, country, genres, sortBy, order)
	if err != nil {
		return nil, 0, helper.NewInternalServerError()
	}
	return novels, totalRows, nil
}
