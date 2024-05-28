package service

import (
	"fmt"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type songService struct {
	repository repository.SongRepository
	log        logger.Logger
}

func NewSongService(repository repository.SongRepository, log logger.Logger) *songService {
	return &songService{repository, log}
}

func (s *songService) CreateSong(userId int, request song.CreateSongRequest) (*models.Song, error) {
	res, err := s.repository.CreateSong(userId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songService) GetSongAll(limit int, offset int, order string) (*[]models.Song, error) {
	res, err := s.repository.GetSongAll(limit, offset, order)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *songService) GetSongBySlug(slug string) (*models.Song, error) {
	res, err := s.repository.GetSongBySlug(slug)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *songService) GetSongById(id int) (*models.Song, error) {
	res, err := s.repository.GetSongById(id)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *songService) FindSongBySlug(slug string, limit int, offset int, order string) (*[]models.Song, error) {
	searchQuery := fmt.Sprintf("%s%s%s", "%", slug, "%")

	res, err := s.repository.FindSongBySlug(searchQuery, limit, offset, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songService) UpdateSong(userId int, songId int, updatedSong song.UpdateSongRequest) (*models.Song, error) {
	res, err := s.repository.UpdateSong(userId, songId, updatedSong)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *songService) DeleteSong(songId int) (*models.Song, error) {
	res, err := s.repository.DeleteSong(songId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
