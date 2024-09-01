package movie

import "github.com/go-playground/validator/v10"

type UpdateMovieRequest struct {
	Title          string `json:"title" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Year           int    `json:"year" validate:"required"`
	PosterImageUrl string `json:"poster" validate:"required"`
	CountryID      int    `json:"country_id" validate:"required"`
}

func (c *UpdateMovieRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
