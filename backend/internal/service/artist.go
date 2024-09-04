package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/artist"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type artistService struct {
	repository repository.ArtistRepository
	log        logger.Logger
}

func NewArtistService(repository repository.ArtistRepository, log logger.Logger) *artistService {
	return &artistService{
		repository: repository,
		log:        log,
	}
}

func (s *artistService) CreateArtist(request artist.CreateArtistRequest) (*models.Artist, error) {
	res, err := s.repository.CreateArtist(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *artistService) GetArtistAll(limit int, offset int, order string) (*[]models.Artist, error) {
	res, err := s.repository.GetArtistAll(limit, offset, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *artistService) GetArtistBySlug(slug string) (*models.Artist, error) {
	res, err := s.repository.GetArtistBySlug(slug)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *artistService) FindArtistsBySlug(slug string, limit int, offset int, order string) (*[]models.Artist, error) {
	res, err := s.repository.FindArtistsBySlug(slug, limit, offset, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *artistService) UpdateArtist(id int, request artist.UpdateArtistRequest) (*models.Artist, error) {
	res, err := s.repository.UpdateArtist(id, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *artistService) DeleteArtistById(id int) (*models.Artist, error) {
	res, err := s.repository.DeleteArtistById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
