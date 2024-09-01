package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type languageService struct {
	repository repository.LanguageRepository
	log        logger.Logger
}

func NewLanguageService(repository repository.LanguageRepository, log logger.Logger) *languageService {
	return &languageService{
		repository: repository,
		log:        log,
	}
}

func (s *languageService) CreateLanguage(language models.Language) (*models.Language, error) {
	res, err := s.repository.CreateLanguage(language)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *languageService) GetLanguageAll() (*[]models.Language, error) {
	res, err := s.repository.GetLanguageAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *languageService) GetLanguageByCode(code string) (*models.Language, error) {
	res, err := s.repository.GetLanguageByCode(code)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *languageService) DeleteLanguageById(id int) (*models.Language, error) {
	res, err := s.repository.DeleteLanguageById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
