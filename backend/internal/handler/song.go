package handler

import (
	"fmt"
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initSongGroup(api *fiber.App) {
	songGroup := api.Group("/api/song")

	songGroup.Get("/", h.handlerGetSongAll)
	songGroup.Get("/slug/:slug", h.handlerGetSongBySlug)

	songGroup.Use(middleware.Protector())
	songGroup.Post("/create", h.handlerCreateSong)
	songGroup.Delete("/delete/:id", h.handlerDeleteSong)
}

// @Summary Get all songs
// @Description Get all songs from application
// @Tags Song
// @Accept json
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
// @Param order_by query string true "Order By"
// @Param search query string false "Search query"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song [get]
func (h *Handler) handlerGetSongAll(c *fiber.Ctx) error {
	pageId, err := strconv.Atoi(c.Query("page_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	orderBy := c.Query("order_by")
	isSupportedOrder := util.IsSupportedOrder(orderBy)
	if !isSupportedOrder {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Unsupported order by argument",
		})
	}

	offset := (pageId - 1) * pageSize
	search := c.Query("search")

	if search != "" {
		res, err := h.service.Song.FindSongBySlug(search, pageSize, offset, orderBy)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.Response{
			Message:    fmt.Sprintf("Successfully retrieve search %s page number %d with %d size", search, pageId, pageSize),
			StatusCode: fiber.StatusOK,
			Data:       res,
		})
	} else {
		songs, err := h.service.Song.GetSongAll(pageSize, offset, orderBy)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.Response{
			Message:    fmt.Sprintf("Successfully retrieve page number %d with %d size", pageId, pageSize),
			StatusCode: fiber.StatusOK,
			Data:       songs,
		})
	}
}

// @Summary Get Song by Slug
// @Description Retrieve a song by its slug
// @Tags Song
// @Produce json
// @Param slug path string true "Song Slug"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song/slug/{slug} [get]
func (h *Handler) handlerGetSongBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.service.Song.GetSongBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve song data by slug",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create Song
// @Description Create a new song for the authenticated user
// @Tags Song
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Song title"
// @Param lyrics formData string true "Song lyrics"
// @Param year formData integer true "Song year"
// @Param album_image formData file true "Song album image"
// @Param country_id formData integer true "Song country id"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song/create [post]
func (h *Handler) handlerCreateSong(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	us := authorization[7:]
	id, err := h.tokenManager.ValidateToken(us)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	countryId, err := strconv.Atoi(c.FormValue("country_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	createReq := song.CreateSongRequest{
		Title:     c.FormValue("title"),
		Lyrics:    c.FormValue("lyrics"),
		Year:      year,
		CountryID: countryId,
	}

	file, err := c.FormFile("album_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s/%s-%d-1", "Song", slug.Make(createReq.Title), year)
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq.AlbumImageUrl = imageUrl

	if err := createReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Song.CreateSong(userId, createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Message:    "Successfully creating song",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete song by ID
// @Description Delete a song by its ID
// @Tags Song
// @Produce json
// @Security BearerAuth
// @Param id path string true "Song ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song/delete/{id} [delete]
func (h *Handler) handlerDeleteSong(c *fiber.Ctx) error {
	songId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	deletedSong, err := h.service.Song.DeleteSong(songId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete song",
		StatusCode: fiber.StatusOK,
		Data:       deletedSong,
	})
}
