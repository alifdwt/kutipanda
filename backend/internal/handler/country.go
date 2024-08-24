package handler

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/country"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCountryGroup(api *fiber.App) {
	country := api.Group("/api/country")

	country.Get("/hello", h.handlerCountryHello)
	country.Get("/", h.handlerCountryAll)
	country.Get("/code/:code", h.handlerCountryByCode)

	country.Use(middleware.Protector())
	country.Post("/create", h.handlerCountryCreate)
}

// @Summary Greet the user for country
// @Description Return a greeting message for country
// @Tags Country
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /country/hello [get]
func (h *Handler) handlerCountryHello(c *fiber.Ctx) error {
	return c.SendString("This is a country handler")
}

// @Summary Get all countries
// @Description Retrieve all countries
// @Tags Country
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /country [get]
func (h *Handler) handlerCountryAll(c *fiber.Ctx) error {
	res, err := h.service.Country.GetCountryAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully get all countries",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get country by code
// @Description Retrieve a country by its code
// @Tags Country
// @Produce json
// @Param code path string true "Country Code"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /country/code/{code} [get]
func (h *Handler) handlerCountryByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	res, err := h.service.Country.GetCountryByCode(code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully get country by code",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create country
// @Description Create a new country
// @Tags Country
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param country body country.CreateCountryRequest true "Request body to create a new country"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /country/create [post]
func (h *Handler) handlerCountryCreate(c *fiber.Ctx) error {
	var body country.CreateCountryRequest
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

	createReq := models.Country{
		Name: body.Name,
		Code: body.Code,
	}

	res, err := h.service.Country.CreateCountry(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully create country",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
