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
	LanguageID    int       `json:"language_id" validate:"required"`
	CountryID     int       `json:"country_id" validate:"required"`
	MovieID       int       `json:"movie_id"`
}

func (c *CreateSongRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
