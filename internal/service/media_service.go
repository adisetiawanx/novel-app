package service

import (
	"fmt"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/google/uuid"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type MediaService interface {
	SaveNovelCover(file io.Reader, fileHeader *multipart.FileHeader) ([]*model.Media, error)
	FindByUrl(url string) (*model.Media, error)
}

type mediaServiceImpl struct {
	repository.MediaRepository
}

func NewMediaService(mediaRepository repository.MediaRepository) MediaService {
	return &mediaServiceImpl{
		MediaRepository: mediaRepository,
	}
}

func (service *mediaServiceImpl) SaveNovelCover(file io.Reader, fileHeader *multipart.FileHeader) ([]*model.Media, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	originalFileName := fileHeader.Filename
	nameWithoutExt := strings.TrimSuffix(originalFileName, path.Ext(originalFileName))

	uniqueID := uuid.New().String()
	timestamp := time.Now().Format("20060102_150405")

	var mediaFiles []*model.Media

	originalWidth := img.Bounds().Dx()

	if originalWidth > 170 {
		resizedImageSmall, err := helper.ResizeImage(img, uint(170))
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		smallFilename := fmt.Sprintf("%s_%s-small.%s", timestamp, uniqueID, format)
		smallPath := fmt.Sprintf("./storage/img/%s", smallFilename)

		err = helper.SaveResizedImage(smallPath, resizedImageSmall, format)
		if err != nil {
			return nil, helper.NewInternalServerError()
		}

		smallFileInfo, err := os.Stat(smallPath)
		if err != nil {
			return nil, helper.NewInternalServerError()
		}

		smallMimeType := fmt.Sprintf("image/%s", format)
		mediaFiles = append(mediaFiles, &model.Media{
			Url:  fmt.Sprintf("https://storage.novelku.id/assets/%s", smallFilename),
			Name: nameWithoutExt + "-small",
			Size: int(smallFileInfo.Size()),
			Type: smallMimeType,
		})
	}

	if originalWidth > 250 {
		resizedImageMedium, err := helper.ResizeImage(img, uint(250))
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		mediumFilename := fmt.Sprintf("%s_%s-medium.%s", timestamp, uniqueID, format)
		mediumPath := fmt.Sprintf("./storage/img/%s", mediumFilename)

		err = helper.SaveResizedImage(mediumPath, resizedImageMedium, format)
		if err != nil {
			return nil, helper.NewInternalServerError()
		}

		mediumFileInfo, err := os.Stat(mediumPath)
		if err != nil {
			return nil, helper.NewInternalServerError()
		}

		mediumMimeType := fmt.Sprintf("image/%s", format)
		mediaFiles = append(mediaFiles, &model.Media{
			Url:  fmt.Sprintf("https://storage.novelku.id/assets/%s", mediumFilename),
			Name: nameWithoutExt + "-medium",
			Size: int(mediumFileInfo.Size()),
			Type: mediumMimeType,
		})
	}

	originalFilename := fmt.Sprintf("%s_%s.%s", timestamp, uniqueID, format)
	originalPath := fmt.Sprintf("./storage/img/%s", originalFilename)

	err = helper.SaveResizedImage(originalPath, img, format)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	originalFileInfo, err := os.Stat(originalPath)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	originalMimeType := fmt.Sprintf("image/%s", format)
	mediaFiles = append(mediaFiles, &model.Media{
		Url:  fmt.Sprintf("https://storage.novelku.id/assets/%s", originalFilename),
		Name: nameWithoutExt,
		Size: int(originalFileInfo.Size()),
		Type: originalMimeType,
	})

	var savedMedia []*model.Media
	for _, media := range mediaFiles {
		newId, err := uuid.NewV7()
		if err != nil {
			return nil, helper.NewInternalServerError()
		}

		media.ID = newId
		savedMediaItem, err := service.MediaRepository.Save(media)
		if err != nil {
			return nil, helper.NewInternalServerError()
		}
		savedMedia = append(savedMedia, savedMediaItem)
	}

	return savedMedia, nil
}

func (service *mediaServiceImpl) FindByUrl(url string) (*model.Media, error) {
	media, err := service.MediaRepository.FindByUrl(url)

	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	if media == nil {
		return nil, helper.NewUserInputError("Media not found")
	}

	return media, nil
}
