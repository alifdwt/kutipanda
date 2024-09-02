package country

import "github.com/go-playground/validator/v10"

type CreateCountryRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required,max=2"`
}

func (c *CreateCountryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
