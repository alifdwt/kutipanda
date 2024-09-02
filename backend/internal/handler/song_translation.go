package handler

import (
	"strconv"

	songtranslation "github.com/alifdwt/kutipanda-backend/internal/domain/requests/song_translation"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/util"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSongTranslationGroup(api *fiber.App) {
	songTranslation := api.Group("/api/song-translation")

	songTranslation.Get("/", h.handlerGetSongTranslationAll)
	// songTranslation.Get("/slug/:slug", h.handlerGetSongTranslationBySlug)

	songTranslation.Use(middleware.Protector())
	songTranslation.Post("/create", h.handlerCreateSongTranslation)
	songTranslation.Put("/update/:id", h.handlerUpdateSongTranslation)
	songTranslation.Delete("/delete/:id", h.handlerDeleteSongTranslation)
}

// @Summary Get all song translation
// @Description Get all song translation
// @Tags Song Translation
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
// @Param order_by query string true "Order By"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song-translation [get]
func (h *Handler) handlerGetSongTranslationAll(c *fiber.Ctx) error {
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

	songTranslations, err := h.service.SongTranslation.GetSongTranslationALl(pageSize, offset, orderBy)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get all song translation",
		StatusCode: fiber.StatusOK,
		Data:       songTranslations,
	})
}

// @Summary Create Song Translation
// @Description Create Song Translation
// @Tags Song Translation
// @Accept json
// @Produce json
// @Param request body songtranslation.CreateSongTranslationRequest true "Create Song Translation"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song-translation/create [post]
func (h *Handler) handlerCreateSongTranslation(c *fiber.Ctx) error {
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

	var body songtranslation.CreateSongTranslationRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.SongTranslation.CreateSongTranslation(userId, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Message:    "Successfully create song translation",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update Song Translation
// @Description Update Song Translation
// @Tags Song Translation
// @Accept json
// @Produce json
// @Param id path int true "Song Translation ID"
// @Param request body songtranslation.UpdateSongTranslationRequest true "Update Song Translation"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song-translation/update/{id} [put]
func (h *Handler) handlerUpdateSongTranslation(c *fiber.Ctx) error {
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

	songTranslationId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body songtranslation.UpdateSongTranslationRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.SongTranslation.UpdateSongTranslationById(userId, songTranslationId, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update song translation",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete Song Translation
// @Description Delete Song Translation
// @Tags Song Translation
// @Accept json
// @Produce json
// @Param id path int true "Song Translation ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /song-translation/delete/{id} [delete]
func (h *Handler) handlerDeleteSongTranslation(c *fiber.Ctx) error {
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

	songTranslationId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.SongTranslation.DeleteSongTranslation(userId, songTranslationId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete song translation",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
