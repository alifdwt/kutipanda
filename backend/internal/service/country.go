package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type countryService struct {
	repository repository.CountryRepository
	log        logger.Logger
}

func NewCountryService(repository repository.CountryRepository, log logger.Logger) *countryService {
	return &countryService{repository, log}
}

func (s *countryService) CreateCountry(country models.Country) (*models.Country, error) {
	res, err := s.repository.CreateCountry(&country)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *countryService) GetCountryAll() (*[]models.Country, error) {
	res, err := s.repository.GetCountryAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *countryService) GetCountryByCode(code string) (*models.Country, error) {
	res, err := s.repository.GetCountryByCode(code)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *countryService) DeleteCountryByCode(code string) (*models.Country, error) {
	res, err := s.repository.DeleteCountryByCode(code)
	if err != nil {
		return nil, err
	}

	return res, nil
}
