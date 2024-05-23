package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/mapper"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type userService struct {
	repository repository.UserRepository
	log        logger.Logger
	mapper     mapper.UserMapping
}

func NewUserService(repository repository.UserRepository, log logger.Logger, mapper mapper.UserMapping) *userService {
	return &userService{repository: repository, log: log, mapper: mapper}
}

func (s *userService) GetUserAll() (*[]responses.UserResponse, error) {
	res, err := s.repository.GetUserAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponses(res)

	return &mapper, nil
}

func (s *userService) GetUserByUsername(username string) (*responses.UserResponse, error) {
	res, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) GetRandomUser(count int) (*[]responses.UserResponse, error) {
	res, err := s.repository.GetRandomUser(count)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponses(res)

	return &mapper, nil
}
