package language

import "github.com/go-playground/validator/v10"

type CreateLanguageRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required,len=2"`
}

func (c *CreateLanguageRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
