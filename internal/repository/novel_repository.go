package repository

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/model"
	"gorm.io/gorm"
	"strings"
)

type NovelRepository interface {
	Save(novel *model.Novel) (*model.Novel, error)
	FindById(novelId string) (*model.Novel, error)
	IsSlugExist(slug string) (bool, error)
	FindAll(page int, pageSize int, title string, status string, country string, genres []string, sortBy string, order string) ([]*model.Novel, int64, error)
}

type novelRepositoryImpl struct {
	DB *gorm.DB
}

func NewNovelRepository(db *gorm.DB) NovelRepository {
	return &novelRepositoryImpl{DB: db}
}

func (repository *novelRepositoryImpl) Save(novel *model.Novel) (*model.Novel, error) {
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(novel.Genres) > 0 {
		var finalGenres []*model.Genre
		for _, genre := range novel.Genres {
			var existingGenre model.Genre
			err := tx.Where("name = ?", genre.Name).First(&existingGenre).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = tx.Create(&genre).Error
					if err != nil {
						tx.Rollback()
						return nil, err
					}
					finalGenres = append(finalGenres, genre)
				} else {
					tx.Rollback()
					return nil, err
				}
			} else {
				finalGenres = append(finalGenres, &existingGenre)
			}
		}
		novel.Genres = finalGenres
	}

	if len(novel.Authors) > 0 {
		var finalAuthors []*model.Author
		for _, author := range novel.Authors {
			var existingAuthor model.Author
			err := tx.Where("name = ?", author.Name).First(&existingAuthor).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = tx.Create(&author).Error
					if err != nil {
						tx.Rollback()
						return nil, err
					}
					finalAuthors = append(finalAuthors, author)
				} else {
					tx.Rollback()
					return nil, err
				}
			} else {
				finalAuthors = append(finalAuthors, &existingAuthor)
			}
		}
		novel.Authors = finalAuthors
	}

	if len(novel.Translators) > 0 {
		var finalTranslators []*model.Translator
		for _, translator := range novel.Translators {
			var existingTranslator model.Translator
			err := tx.Where("name = ?", translator.Name).First(&existingTranslator).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = tx.Create(&translator).Error
					if err != nil {
						tx.Rollback()
						return nil, err
					}
					finalTranslators = append(finalTranslators, translator)
				} else {
					tx.Rollback()
					return nil, err
				}
			} else {
				finalTranslators = append(finalTranslators, &existingTranslator)
			}
		}
		novel.Translators = finalTranslators
	}

	if err := tx.Create(&novel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return novel, nil
}

func (repository *novelRepositoryImpl) FindById(novelId string) (*model.Novel, error) {
	var novel model.Novel
	err := repository.DB.Where("id = ?", novelId).First(&novel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &novel, nil
}

func (repository *novelRepositoryImpl) FindAll(page int, pageSize int, title string, status string, country string, genres []string, sortBy string, order string) ([]*model.Novel, int64, error) {
	var novels []*model.Novel
	var totalRows int64

	// Start query without preload
	query := repository.DB.Model(&model.Novel{})

	//take 2 latest chapters
	query = query.Preload("Chapters", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, novel_id, title, slug").Order("created_at DESC").Limit(2)
	})

	// Filters
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if country != "" {
		query = query.Where("country = ?", country)
	}
	if len(genres) > 0 {
		query = query.
			Joins("JOIN novelku.novel_genres ng ON ng.novel_id = novels.id").
			Joins("JOIN novelku.genres g ON g.id = ng.genre_id").
			Where("g.name IN ?", genres).
			Group("novels.id")
	}

	// Count total matching rows
	err := query.Count(&totalRows).Error
	if err != nil {
		return nil, 0, err
	}

	// Handle dynamic sort
	validSortFields := map[string]string{
		"views":      "view_total",
		"bookmarks":  "bookmark_total",
		"rating":     "rating_total", // handled separately
		"created_at": "created_at",
		"title":      "title",
		"country":    "country",
		"status":     "status",
	}

	if field, ok := validSortFields[sortBy]; ok {
		if sortBy == "rating" {
			if order == "asc" {
				query = query.Order("rating_total ASC, vote_total ASC")
			} else {
				query = query.Order("rating_total DESC, vote_total DESC")
			}
		} else {
			query = query.Order(field + " " + strings.ToUpper(order))
		}
	} else {
		query = query.Order("created_at DESC") // Default
	}

	// Pagination
	offset := (page - 1) * pageSize
	err = query.
		Limit(pageSize).
		Offset(offset).
		Find(&novels).Error
	if err != nil {
		return nil, 0, err
	}

	return novels, totalRows, nil
}

func (repository *novelRepositoryImpl) IsSlugExist(slug string) (bool, error) {
	var novel model.Novel
	err := repository.DB.Where("slug = ?", slug).First(&novel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
