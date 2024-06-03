package songtranslation

import "github.com/go-playground/validator/v10"

type UpdateSongTranslationRequest struct {
	SongID           int    `json:"song_id" validate:"required"`
	TranslatedTitle  string `json:"translated_title" validate:"required"`
	TranslatedLyrics string `json:"translated_lyrics" validate:"required"`
	Language         string `json:"language" validate:"required"`
}

func (u *UpdateSongTranslationRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
