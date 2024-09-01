package handler

import (
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/language"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initLanguageGroup(api *fiber.App) {
	language := api.Group("/api/language")

	language.Get("/", h.handlerGetLanguageAll)
	language.Get("/code/:code", h.handlerGetLanguageByCode)

	language.Use(middleware.Protector())
	language.Post("/create", h.handlerCreateLanguage)
	language.Delete("/delete/:id", h.handlerDeleteLanguageById)
}

// @Summary Get all language
// @Description Get all language
// @Tags Language
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /language [get]
func (h *Handler) handlerGetLanguageAll(c *fiber.Ctx) error {
	languages, err := h.service.Language.GetLanguageAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success get all language",
		StatusCode: fiber.StatusOK,
		Data:       languages,
	})
}

// @Summary Get language by code
// @Description Get language by code
// @Tags Language
// @Produce json
// @Param code path string true "Language Code"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /language/code/{code} [get]
func (h *Handler) handlerGetLanguageByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	language, err := h.service.Language.GetLanguageByCode(code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success get language by code",
		StatusCode: fiber.StatusOK,
		Data:       language,
	})
}

// @Summary Create language
// @Description Create language
// @Tags Language
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param language body language.CreateLanguageRequest true "Create language"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /language/create [post]
func (h *Handler) handlerCreateLanguage(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	us := authorization[7:]
	_, err := h.tokenManager.ValidateToken(us)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	var body language.CreateLanguageRequest
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

	createReq := models.Language{
		Name: body.Name,
		Code: body.Code,
	}

	res, err := h.service.Language.CreateLanguage(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success create language",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete language by id
// @Description Delete language by id
// @Tags Language
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Language ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /language/delete/{id} [delete]
func (h *Handler) handlerDeleteLanguageById(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	us := authorization[7:]
	_, err := h.tokenManager.ValidateToken(us)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Language.DeleteLanguageById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success delete language",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
