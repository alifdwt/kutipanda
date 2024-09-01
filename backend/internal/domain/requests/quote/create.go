package quote

import "github.com/go-playground/validator/v10"

type CreateQuoteRequest struct {
	QuoteText            string `json:"quote_text" validate:"required"`
	QuoteTransliteration string `json:"quote_transliteration"`
	MovieID              int    `json:"movie_id" validate:"required"`
	LanguageID           int    `json:"language_id" validate:"required"`
}

func (c *CreateQuoteRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
