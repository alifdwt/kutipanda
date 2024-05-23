package repository

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/user"
	"github.com/alifdwt/kutipanda-backend/internal/models"
)

type UserRepository interface {
	GetUserAll() (*[]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(registerReq *auth.RegisterRequest) (*models.User, error)
	UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
	GetRandomUser(count int) (*[]models.User, error)
}

type MovieRepository interface {
	CreateMovie(movie movie.CreateMovieRequest) (*models.Movie, error)
	GetMovieAll() (*[]models.Movie, error)
	GetMovieBySlug(slug string) (*models.Movie, error)
	GetMovieById(id int) (*models.Movie, error)
	UpdateMovieById(id int, updatedMovie movie.UpdateMovieRequest) (*models.Movie, error)
	DeleteMovie(id int) (*models.Movie, error)
}
