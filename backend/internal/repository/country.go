package repository

import (
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"gorm.io/gorm"
)

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *countryRepository {
	return &countryRepository{db}
}

func (r *countryRepository) CreateCountry(country *models.Country) (*models.Country, error) {
	var countryModel models.Country

	db := r.db.Model(&countryModel)

	if err := db.Debug().Create(&country).Error; err != nil {
		return nil, err
	}

	return country, nil
}

func (r *countryRepository) GetCountryAll() (*[]models.Country, error) {
	var countries []models.Country

	db := r.db.Model(&countries)

	if err := db.Debug().Find(&countries).Error; err != nil {
		return nil, err
	}

	return &countries, nil
}

func (r *countryRepository) GetCountryByCode(code string) (*models.Country, error) {
	var country models.Country

	db := r.db.Model(&country)

	if err := db.Debug().Where("code = ?", code).First(&country).Error; err != nil {
		return nil, err
	}

	return &country, nil
}

func (r *countryRepository) DeleteCountryByCode(code string) (*models.Country, error) {
	var country models.Country

	db := r.db.Model(&country)

	if err := db.Debug().Where("code = ?", code).Delete(&country).Error; err != nil {
		return nil, err
	}

	return &country, nil
}
