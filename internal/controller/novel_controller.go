package controller

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/dto"
	"github.com/adisetiawanx/novel-app/internal/dto/request"
	"github.com/adisetiawanx/novel-app/internal/dto/response"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

type NovelController interface {
	Create(context echo.Context) error
	FindAll(context echo.Context) error
}

type novelControllerImpl struct {
	service.NovelService
}

func NewNovelController(novelService service.NovelService) NovelController {
	return &novelControllerImpl{
		NovelService: novelService,
	}
}

func (controller *novelControllerImpl) Create(ctx echo.Context) error {
	var req request.NovelCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	var baseError *helper.BaseError
	novel, err := controller.NovelService.Save(&model.Novel{
		PostStatus:       req.PostStatus,
		Title:            req.Title,
		Slug:             req.Slug,
		AlternativeTitle: req.AlternativeTitle,
		Synopsis:         req.Synopsis,
		Status:           req.Status,
		ReleaseYear:      req.ReleaseYear,
		Country:          req.Country,
		Genres:           req.Genres,
		Authors:          req.Authors,
		Translators:      req.Translators,
	})
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, dto.APIResponse{
		Message: "Novel Created Successfully",
		Data: response.NovelCreateResponse{
			Id:     novel.ID.String(),
			Title:  novel.Title,
			Slug:   novel.Slug,
			Status: novel.Status,
		},
	})
}

func (controller *novelControllerImpl) FindAll(ctx echo.Context) error {
	var novelPerPage = 10

	// Parsing query parameters
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(ctx.QueryParam("page_size"))
	if pageSize == 0 {
		pageSize = novelPerPage
	}

	title := ctx.QueryParam("title")
	status := ctx.QueryParam("status")
	country := ctx.QueryParam("country")
	genresParam := ctx.QueryParam("genres")
	sortBy := ctx.QueryParam("sort_by")
	if sortBy == "" {
		sortBy = "created_at" // Default sort
	}
	order := ctx.QueryParam("order")
	if order == "" {
		order = "desc" // Default order
	}

	// Parse genres, if provided
	var genres []string
	if genresParam != "" {
		genres = strings.Split(genresParam, ",")
	}

	// Call the service method to fetch novels with filters
	var baseError *helper.BaseError
	novels, total, err := controller.NovelService.FindAll(page, pageSize, title, status, country, genres, sortBy, order)
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	// Return the result
	return ctx.JSON(http.StatusOK, dto.APIResponse{
		Message: "Novel Fetched Successfully",
		Data: response.NovelsGetResponseWrapper{
			Novels:   helper.ToNovelsGetResponse(novels),
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	})
}
