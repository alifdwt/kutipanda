package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/mapper"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/pkg/auth"
	"github.com/alifdwt/kutipanda-backend/pkg/hashing"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
)

type Service struct {
	Auth  AuthService
	User  UserService
	Movie MovieService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:  NewAuthService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Token, deps.Mapper.UserMapper),
		User:  NewUserService(deps.Repository.User, deps.Logger, deps.Mapper.UserMapper),
		Movie: NewMovieService(deps.Repository.Movie, deps.Logger),
	}
}
