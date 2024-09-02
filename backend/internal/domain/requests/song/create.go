package song

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateSongRequest struct {
	Title         string    `json:"title" validate:"required"`
	Lyrics        string    `json:"lyrics" validate:"required"`
	ReleaseDate   time.Time `json:"release_date" validate:"required"`
	AlbumImageUrl string    `json:"album_image_url" validate:"required,url"`
	CountryID     int       `json:"country_id" validate:"required"`
}

func (c *CreateSongRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
