package handler

import (
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/quote"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/util"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initQuoteGroup(api *fiber.App) {
	quote := api.Group("/api/quote")

	quote.Get("/hello", h.handlerQuoteHello)
	quote.Get("/", h.handlerQuoteAll)
	quote.Get("/random/:count", h.handlerQuoteRandom)

	quote.Use(middleware.Protector())
	quote.Post("/create", h.handlerQuoteCreate)
	quote.Put("/update/:id", h.handlerQuoteUpdate)
	quote.Delete("/delete/:id", h.handlerQuoteDelete)
}

// @Summary Greet hello
// @Description Greet hello
// @Tags Quote
// @Produce plain
// @Success 200 {string} string "Hello from kutipanda"
// @Router /quote/hello [get]
func (h *Handler) handlerQuoteHello(c *fiber.Ctx) error {
	return c.SendString("Hello from kutipanda")
}

// @Summary Get all quotes
// @Description Get all quotes
// @Tags Quote
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
// @Param order_by query string true "Order By"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /quote [get]
func (h *Handler) handlerQuoteAll(c *fiber.Ctx) error {
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

	res, err := h.service.Quote.GetQuoteAll(pageSize, offset, orderBy)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success get all quotes",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get random quotes
// @Description Get random quotes
// @Tags Quote
// @Produce json
// @Param count path int true "Count"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /quote/random/{count} [get]
func (h *Handler) handlerQuoteRandom(c *fiber.Ctx) error {
	count, err := c.ParamsInt("count")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Quote.GetRandomQuote(count)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success get random quotes",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create quote
// @Description Create quote
// @Tags Quote
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body quote.CreateQuoteRequest true "Create quote"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /quote/create [post]
func (h *Handler) handlerQuoteCreate(c *fiber.Ctx) error {
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

	var body quote.CreateQuoteRequest
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

	res, err := h.service.Quote.CreateQuote(userId, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Message:    "Success create quote",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update quote
// @Description Update quote
// @Tags Quote
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Quote ID"
// @Param request body quote.UpdateQuoteRequest true "Update quote"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /quote/update/{id} [put]
func (h *Handler) handlerQuoteUpdate(c *fiber.Ctx) error {
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

	quoteId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body quote.UpdateQuoteRequest
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

	res, err := h.service.Quote.UpdateQuote(userId, quoteId, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success update quote",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete quote
// @Description Delete quote
// @Tags Quote
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Quote ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /quote/delete/{id} [delete]
func (h *Handler) handlerQuoteDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Quote.DeleteQuoteById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Success delete quote",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
