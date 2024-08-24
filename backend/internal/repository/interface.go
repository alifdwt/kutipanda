package repository

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	songtranslation "github.com/alifdwt/kutipanda-backend/internal/domain/requests/song_translation"
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

type SongRepository interface {
	CreateSong(userId int, request song.CreateSongRequest) (*models.Song, error)
	GetSongAll(limit int, offset int, order string) (*[]models.Song, error)
	GetSongBySlug(slug string) (*models.Song, error)
	GetSongById(id int) (*models.Song, error)
	FindSongBySlug(slug string, limit int, offset int, order string) (*[]models.Song, error)
	UpdateSong(userId int, songId int, updatedSong song.UpdateSongRequest) (*models.Song, error)
	DeleteSong(songId int) (*models.Song, error)
}

type SongTranslationRepository interface {
	CreateSongTranslation(userId int, request songtranslation.CreateSongTranslationRequest) (*models.SongTranslation, error)
	GetSongTranslationAll(limit int, offset int, order string) (*[]models.SongTranslation, error)
	GetSongTranslationById(id int) (*models.SongTranslation, error)
	GetSongTranslationBySlug(slug string) (*models.SongTranslation, error)
	UpdateSongTranslation(userId int, songTranslationId int, updatedSongTranslation songtranslation.UpdateSongTranslationRequest) (*models.SongTranslation, error)
	DeleteSongTranslation(userId int, songTranslationId int) (*models.SongTranslation, error)
}

type CountryRepository interface {
	CreateCountry(country *models.Country) (*models.Country, error)
	GetCountryAll() (*[]models.Country, error)
	GetCountryByCode(code string) (*models.Country, error)
	DeleteCountryByCode(code string) (*models.Country, error)
}
