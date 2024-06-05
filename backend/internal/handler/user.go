package handler

import (
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initUserGroup(api *fiber.App) {
	user := api.Group("/api/user")

	user.Get("/", h.handlerGetUserAll)
	user.Get("/username/:username", h.handlerGetUserByUsername)

	user.Use(middleware.Protector())
	user.Get("/me", h.handlerGetUserMe)
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user [get]
func (h *Handler) handlerGetUserAll(c *fiber.Ctx) error {
	res, err := h.service.User.GetUserAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get all user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get user data by username
// @Description Retrieve user data by username
// @Tags User
// @Param username path string true "User username"
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/username/{username} [get]
func (h *Handler) handlerGetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	res, err := h.service.User.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get user data by username",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get user data
// @Description Retrieve logged in user data
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /user/me [get]
func (h *Handler) handlerGetUserMe(c *fiber.Ctx) error {
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

	res, err := h.service.User.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get user data",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
