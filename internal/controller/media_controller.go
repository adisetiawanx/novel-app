package controller

import (
	"github.com/adisetiawanx/novel-app/internal/dto"
	"github.com/adisetiawanx/novel-app/internal/dto/response"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

const (
	MaxMemory   = 10 << 20        // 10MB
	MaxFileSize = 3 * 1024 * 1024 // 3MB
)

type MediaController interface {
	UploadNovelCover(context echo.Context) error
}

type mediaControllerImpl struct {
	service.MediaService
}

func NewMediaController(mediaService service.MediaService) MediaController {
	return &mediaControllerImpl{
		MediaService: mediaService,
	}
}

func (controller *mediaControllerImpl) UploadNovelCover(ctx echo.Context) error {
	err := ctx.Request().ParseMultipartForm(MaxMemory)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "Failed to parse form data",
			Data:    nil,
		})
	}

	file, fileHeader, err := ctx.Request().FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.APIResponse{
			Message: "Failed to retrieve file from form data",
			Data:    nil,
		})
	}
	defer file.Close()

	if fileHeader.Size > MaxFileSize {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "File too large. Max 3MB allowed",
			Data:    nil,
		})
	}

	var baseError *helper.BaseError
	images, err := controller.MediaService.SaveNovelCover(file, fileHeader)
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, dto.APIResponse{
		Message: "Image uploaded successfully",
		Data: response.MediaUploadNovelCoverResponseWrapper{
			Media: helper.ToMediaUploadNovelCoverResponse(images),
			Total: len(images),
		},
	})
}
