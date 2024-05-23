package handler

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initUserGroup(api *fiber.App) {
	user := api.Group("/api/user")

	user.Get("/", h.handlerGetUserAll)
	user.Get("/:username", h.handlerGetUserByUsername)
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
// @Router /user/{username} [get]
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
