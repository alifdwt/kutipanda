package songtranslation

import "github.com/go-playground/validator/v10"

type CreateSongTranslationRequest struct {
	SongID           int    `json:"song_id" validate:"required"`
	TranslatedTitle  string `json:"translated_title" validate:"required"`
	TranslatedLyrics string `json:"translated_lyrics" validate:"required"`
	Language         string `json:"language" validate:"required"`
}

func (c *CreateSongTranslationRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
