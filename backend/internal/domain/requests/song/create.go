package song

import "github.com/go-playground/validator/v10"

type CreateSongRequest struct {
	Title         string `json:"title" validate:"required"`
	Lyrics        string `json:"lyrics" validate:"required"`
	Year          int    `json:"year" validate:"required"`
	AlbumImageUrl string `json:"album_image_url" validate:"required,url"`
	Language      string `json:"language" validate:"required"`
}

func (c *CreateSongRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
