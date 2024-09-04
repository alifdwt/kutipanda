package handler

import (
	"fmt"
	"strconv"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/artist"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/middleware"
	"github.com/alifdwt/kutipanda-backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (h *Handler) initArtistGroup(api *fiber.App) {
	artist := api.Group("/api/artist")

	artist.Get("/", h.handlerArtistAll)
	artist.Get("/slug/:slug", h.handlerGetArtistBySlug)

	artist.Use(middleware.Protector())
	artist.Post("/create", h.handlerArtistCreate)
	artist.Put("/update/:id", h.handlerArtistUpdate)
	artist.Delete("/delete/:id", h.handlerArtistDelete)
}

// @Summary Get all artists
// @Description Get all artists
// @Tags Artist
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
// @Param order_by query string true "Order By"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Route /artist [get]
func (h *Handler) handlerArtistAll(c *fiber.Ctx) error {
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

	res, err := h.service.Artist.GetArtistAll(pageSize, offset, orderBy)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get all artists",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get artist by slug
// @Description Get artist by slug
// @Tags Artist
// @Produce json
// @Param slug path string true "Slug"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Failure 404 {object} responses.ErrorMessage
// @Route /artist/slug/{slug} [get]
func (h *Handler) handlerGetArtistBySlug(c *fiber.Ctx) error {
	artistSlug := c.Params("slug")

	res, err := h.service.Artist.GetArtistBySlug(artistSlug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully get artist by slug",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create artist
// @Description Create new artist
// @Tags Artist
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param sex formData string true "Sex"
// @Param country_id formData integer true "Country ID"
// @Param image_url formData file true "Artist Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /artist/create [post]
func (h *Handler) handlerArtistCreate(c *fiber.Ctx) error {
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

	countryId, err := strconv.Atoi(c.FormValue("country_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	createReq := artist.CreateArtistRequest{
		Name:      c.FormValue("name"),
		Sex:       c.FormValue("sex"),
		CountryID: countryId,
	}

	file, err := c.FormFile("image_url")
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

	fileName := fmt.Sprintf("%s/%s", "Artist", slug.Make(createReq.Name))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq.ImageURL = imageUrl

	if err := createReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Artist.CreateArtist(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Message:    "Successfully create artist",
		StatusCode: fiber.StatusCreated,
		Data:       res,
	})
}

// @Summary Update artist
// @Description Update artist
// @Tags Artist
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param id formData integer true "Artist ID"
// @Param name formData string true "Name"
// @Param sex formData string true "Sex"
// @Param country_id formData integer true "Country ID"
// @Param image_url formData file true "Artist Image"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /artist/update/{id} [put]
func (h *Handler) handlerArtistUpdate(c *fiber.Ctx) error {
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

	artistId, err := strconv.Atoi(c.Params("id"))
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

	updateReq := artist.UpdateArtistRequest{
		Name:      c.FormValue("name"),
		Sex:       c.FormValue("sex"),
		CountryID: countryId,
	}

	file, err := c.FormFile("image_url")
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

	fileName := fmt.Sprintf("%s/%s", "Artist", slug.Make(updateReq.Name))
	imageUrl, err := h.cloudinary.UploadToCloudinary(uploadedFile, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq.ImageURL = imageUrl

	if err := updateReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.service.Artist.UpdateArtist(artistId, updateReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully update artist",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete artist
// @Description Delete artist
// @Tags Artist
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Artist ID"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.ErrorMessage
// @Router /artist/delete/{id} [delete]
func (h *Handler) handlerArtistDelete(c *fiber.Ctx) error {
	artistId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	deletedSong, err := h.service.Artist.DeleteArtistById(artistId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Message:    "Successfully delete artist",
		StatusCode: fiber.StatusOK,
		Data:       deletedSong,
	})
}
