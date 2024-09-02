package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initMovieGroup(api *fiber.App) {
	movie := api.Group("/api/movie")

	movie.Get("/hello", h.handlerMovieHello)
	movie.Get("/", h.handlerMovieAll)
	movie.Get("/slug/:slug", h.handlerMovieBySlug)

	movie.Use(middleware.Protector())
	movie.Get("/:id", h.handlerMovieById)
	movie.Post("/create", h.handlerMovieCreate)
}

// @Summary Greet the user for movie
// @Description Return a greeting message for movie
// @Tags Movie
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /movie/hello [get]
func (h *Handler) handlerMovieHello(c *fiber.Ctx) error {
	return c.SendString("This is a movie handler")
}

// @Summary Get all movies
// @Description Retrieve all movies
// @Tags Movie
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /movie [get]
func (h *Handler) handlerMovieAll(c *fiber.Ctx) error {
	res, err := h.service.Movie.GetMovieAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(responses.Response{
		Message:    "Successfully retrieve all movies",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get movie by slug
// @Description Retrieve a movie by its slug
// @Tags Movie
// @Produce json
// @Param slug path string true "Movie Slug"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /movie/slug/{slug} [get]
func (h *Handler) handlerMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.service.Movie.GetMovieBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve movie data by slug",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get movie by ID
// @Description Retrieve movie data by its ID
// @Tags Movie
// @Produce json
// @Security BearerAuth
// @Param id path int true "Movie ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /movie/{id} [get]
func (h *Handler) handlerMovieById(c *fiber.Ctx) error {
	movieIdStr := c.Params("id")
	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Movie.GetMovieById(movieId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully retrieve movie data by id",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create movie
// @Descpription Create a new movie
// @Tags Movie
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Movie title"
// @Param description formData string true "Movie description"
// @Param release_date formData string true "Movie release date"
// @Param poster_image formData file true "Movie poster image"
// @Param country_id formData integer true "Movie country id"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /movie/create [post]
func (h *Handler) handlerMovieCreate(c *fiber.Ctx) error {
	countryId, err := strconv.Atoi(c.FormValue("country_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	releaseDate, err := time.Parse("2006-01-02", c.FormValue("release_date"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	createReq := movie.CreateMovieRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		ReleaseDate: releaseDate,
		CountryID:   countryId,
	}

	file, err := c.FormFile("poster_image")
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

	fileName := fmt.Sprintf("%s/%s-%d-1", "Movie", slug.Make(createReq.Title), releaseDate.Year())
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq.PosterImageUrl = imageUrl

	if err := createReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Movie.CreateMovie(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Message:    "Successfully creating movie",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
