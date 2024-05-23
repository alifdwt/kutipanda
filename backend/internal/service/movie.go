package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type movieService struct {
	repository repository.MovieRepository
	log        logger.Logger
}

func NewMovieService(repository repository.MovieRepository, log logger.Logger) *movieService {
	return &movieService{repository, log}
}

func (s *movieService) CreateMovie(movie movie.CreateMovieRequest) (*models.Movie, error) {
	res, err := s.repository.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) GetMovieAll() (*[]models.Movie, error) {
	res, err := s.repository.GetMovieAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) GetMovieBySlug(slug string) (*models.Movie, error) {
	res, err := s.repository.GetMovieBySlug(slug)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) GetMovieById(id int) (*models.Movie, error) {
	res, err := s.repository.GetMovieById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) UpdateMovieById(id int, updatedMovie movie.UpdateMovieRequest) (*models.Movie, error) {
	res, err := s.repository.UpdateMovieById(id, updatedMovie)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) DeleteMovie(id int) (*models.Movie, error) {
	res, err := s.repository.DeleteMovie(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
