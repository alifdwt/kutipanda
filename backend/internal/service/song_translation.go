package service

import (
	songtranslation "github.com/alifdwt/kutipanda-backend/internal/domain/requests/song_translation"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type songTranslationService struct {
	repository repository.SongTranslationRepository
	log        logger.Logger
}

func NewSongTranslationService(repository repository.SongTranslationRepository, log logger.Logger) *songTranslationService {
	return &songTranslationService{repository, log}
}

func (s *songTranslationService) CreateSongTranslation(userId int, request songtranslation.CreateSongTranslationRequest) (*models.SongTranslation, error) {
	res, err := s.repository.CreateSongTranslation(userId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songTranslationService) GetSongTranslationALl(limit int, offset int, order string) (*[]models.SongTranslation, error) {
	res, err := s.repository.GetSongTranslationAll(limit, offset, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songTranslationService) GetSongTranslationById(id int) (*models.SongTranslation, error) {
	res, err := s.repository.GetSongTranslationById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songTranslationService) GetSongTranslationBySlug(slug string) (*models.SongTranslation, error) {
	res, err := s.repository.GetSongTranslationBySlug(slug)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songTranslationService) UpdateSongTranslationById(userId int, songTranslationId int, updatedSongTranslation songtranslation.UpdateSongTranslationRequest) (*models.SongTranslation, error) {
	res, err := s.repository.UpdateSongTranslation(userId, songTranslationId, updatedSongTranslation)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songTranslationService) DeleteSongTranslation(userId int, songTranslationId int) (*models.SongTranslation, error) {
	res, err := s.repository.DeleteSongTranslation(userId, songTranslationId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
