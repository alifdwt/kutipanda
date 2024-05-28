package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	"github.com/alifdwt/kutipanda-backend/internal/domain/responses"
	"github.com/alifdwt/kutipanda-backend/internal/models"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*responses.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
	RefreshToken(req auth.RefreshTokenRequest) (*responses.Token, error)
}

type UserService interface {
	GetUserAll() (*[]responses.UserResponse, error)
	GetUserByUsername(username string) (*responses.UserResponse, error)
	GetRandomUser(count int) (*[]responses.UserResponse, error)
}

type MovieService interface {
	CreateMovie(movie movie.CreateMovieRequest) (*models.Movie, error)
	GetMovieAll() (*[]models.Movie, error)
	GetMovieBySlug(slug string) (*models.Movie, error)
	GetMovieById(id int) (*models.Movie, error)
	UpdateMovieById(id int, updatedMovie movie.UpdateMovieRequest) (*models.Movie, error)
	DeleteMovie(id int) (*models.Movie, error)
}

type SongService interface {
	CreateSong(userId int, request song.CreateSongRequest) (*models.Song, error)
	GetSongAll(limit int, offset int, order string) (*[]models.Song, error)
	GetSongBySlug(slug string) (*models.Song, error)
	GetSongById(id int) (*models.Song, error)
	FindSongBySlug(slug string, limit int, offset int, order string) (*[]models.Song, error)
	UpdateSong(userId int, songId int, updatedSong song.UpdateSongRequest) (*models.Song, error)
	DeleteSong(songId int) (*models.Song, error)
}
