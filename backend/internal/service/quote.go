package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/quote"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type quoteService struct {
	repository repository.QuoteRepository
	log        logger.Logger
}

func NewQuoteService(repository repository.QuoteRepository, log logger.Logger) *quoteService {
	return &quoteService{repository, log}
}

func (s *quoteService) CreateQuote(userId int, quote quote.CreateQuoteRequest) (*models.Quote, error) {
	res, err := s.repository.CreateQuote(userId, quote)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *quoteService) GetQuoteAll() (*[]models.Quote, error) {
	res, err := s.repository.GetQuoteAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *quoteService) GetRandomQuote(count int) (*[]models.Quote, error) {
	res, err := s.repository.GetRandomQuote(count)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *quoteService) DeleteQuoteById(id int) (*models.Quote, error) {
	res, err := s.repository.DeleteQuoteById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
