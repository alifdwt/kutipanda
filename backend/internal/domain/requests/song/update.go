package song

import "github.com/go-playground/validator/v10"

type UpdateSongRequest struct {
	Title         string `json:"title" validate:"required"`
	Lyrics        string `json:"lyrics" validate:"required"`
	Year          int    `json:"year" validate:"required"`
	AlbumImageUrl string `json:"album_image_url" validate:"required,url"`
	CountryID     int    `json:"country_id" validate:"required"`
}

func (c *UpdateSongRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return nil
	}

	return nil
}
