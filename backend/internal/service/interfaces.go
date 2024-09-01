package service

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/auth"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/quote"
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/song"
	songtranslation "github.com/alifdwt/kutipanda-backend/internal/domain/requests/song_translation"
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
	GetUserById(id int) (*responses.UserResponse, error)
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

type SongTranslationService interface {
	CreateSongTranslation(userId int, request songtranslation.CreateSongTranslationRequest) (*models.SongTranslation, error)
	GetSongTranslationALl(limit int, offset int, order string) (*[]models.SongTranslation, error)
	GetSongTranslationById(id int) (*models.SongTranslation, error)
	GetSongTranslationBySlug(slug string) (*models.SongTranslation, error)
	UpdateSongTranslationById(userId int, songTranslationId int, updatedSongTranslation songtranslation.UpdateSongTranslationRequest) (*models.SongTranslation, error)
	DeleteSongTranslation(userId int, songTranslationId int) (*models.SongTranslation, error)
}

type CountryService interface {
	CreateCountry(country models.Country) (*models.Country, error)
	GetCountryAll() (*[]models.Country, error)
	GetCountryByCode(code string) (*models.Country, error)
	DeleteCountryByCode(code string) (*models.Country, error)
}

type QuoteService interface {
	CreateQuote(userId int, quote quote.CreateQuoteRequest) (*models.Quote, error)
	GetQuoteAll() (*[]models.Quote, error)
	GetRandomQuote(count int) (*[]models.Quote, error)
	DeleteQuoteById(id int) (*models.Quote, error)
}
