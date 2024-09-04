package artist

import "github.com/go-playground/validator/v10"

type CreateArtistRequest struct {
	Name      string `json:"name" validate:"required"`
	Sex       string `json:"sex" validate:"required,oneof=male female"`
	CountryID int    `json:"country_id" validate:"required"`
	ImageURL  string `json:"image_url" validate:"required,url"`
}

func (c *CreateArtistRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
