package handler

import (
	"fmt"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initAuthGroup(api *fiber.App) {
	auth := api.Group("/api/auth")

	auth.Get("/hello", h.handlerHello)
	auth.Post("/register", h.register)
	auth.Post("/login", h.login)
}

// @Summary Greet the user
// @Description Return a greeting message to the user
// @Tags Auth
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /auth/hello [get]
func (h *Handler) handlerHello(c *fiber.Ctx) error {
	return c.SendString("This is auth handler")
}

// @Summary Register to Kutipanda
// @Description Create a new user
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Email"
// @Param username formData string true "Username"
// @Param full_name formData string true "Full Name"
// @Param password formData string true "Password"
// @Param description formData string false "Description"
// @Param profile_image formData file true "Profile Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/register [post]
func (h *Handler) register(c *fiber.Ctx) error {
	registerReq := auth.RegisterRequest{
		Email:       c.FormValue("email"),
		Username:    c.FormValue("username"),
		FullName:    c.FormValue("full_name"),
		Password:    c.FormValue("password"),
		Description: c.FormValue("description"),
	}

	file, err := c.FormFile("profile_image")
	if err == nil {
		c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Profile image is required",
			Error:      true,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Error:      true,
		})
	}
	defer uploadedFile.Close()

	fileName := fmt.Sprintf("%s/%s-1", "User", slug.Make(registerReq.FullName))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Error:      true,
		})
	}

	registerReq.ProfileImageURL = imageUrl

	if err := registerReq.Validate(); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	user, err := h.service.Auth.Register(&registerReq)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully register user",
		StatusCode: fiber.StatusOK,
		Data:       user,
	})
}

// @Summary Login to Kutipanda
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User credentials"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /auth/login [post]
func (h *Handler) login(c *fiber.Ctx) error {
	var body auth.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	res, err := h.service.Auth.Login(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Login success",
		Data:       res,
	})
}
