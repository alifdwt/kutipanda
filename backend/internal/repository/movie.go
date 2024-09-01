package repository

import (
	"fmt"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/movie"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) CreateMovie(movie movie.CreateMovieRequest) (*models.Movie, error) {
	var movieModel models.Movie

	db := r.db.Model(&movieModel)

	movieModel.Title = movie.Title
	movieModel.Description = movie.Description
	movieModel.Year = movie.Year
	movieModel.PosterImageUrl = movie.PosterImageUrl
	movieModel.CountryID = movie.CountryID

	movieModel.Slug = fmt.Sprintf("%s-%d", slug.Make(movie.Title), movie.Year)

	result := db.Debug().Create(&movieModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &movieModel, nil
}

func (r *movieRepository) GetMovieAll() (*[]models.Movie, error) {
	var movies []models.Movie

	db := r.db.Model(&movies)

	result := db.Debug().Preload("Country").Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}

	return &movies, nil
}

func (r *movieRepository) GetMovieBySlug(slug string) (*models.Movie, error) {
	var movie models.Movie

	db := r.db.Model(&movie)

	checkMovieBySlug := db.Debug().Preload("Country").Preload("Actors").Preload("Songs").Where("slug = ?", slug).First(&movie)
	if checkMovieBySlug.Error != nil {
		return nil, checkMovieBySlug.Error
	}

	return &movie, nil
}

func (r *movieRepository) GetMovieById(id int) (*models.Movie, error) {
	var movie models.Movie

	db := r.db.Model(&movie)

	checkMovieById := db.Debug().Where("id = ?", id).First(&movie)
	if checkMovieById.Error != nil {
		return nil, checkMovieById.Error
	}

	return &movie, nil
}

func (r *movieRepository) UpdateMovieById(id int, updatedMovie movie.UpdateMovieRequest) (*models.Movie, error) {
	var movie models.Movie

	db := r.db.Model(&movie)

	checkMovieById := db.Debug().Where("id = ?", id).First(&movie)
	if checkMovieById.Error != nil {
		return nil, checkMovieById.Error
	}

	movie.Title = updatedMovie.Title
	movie.Description = updatedMovie.Description
	movie.Year = updatedMovie.Year
	movie.PosterImageUrl = updatedMovie.PosterImageUrl
	movie.CountryID = updatedMovie.CountryID

	updateMovie := db.Debug().Updates(&movie)
	if updateMovie.Error != nil {
		return nil, updateMovie.Error
	}

	return &movie, nil
}

func (r *movieRepository) DeleteMovie(id int) (*models.Movie, error) {
	var movie models.Movie

	db := r.db.Model(&movie)

	checkMovieById := db.Debug().Where("id = ?", id).First(&movie)
	if checkMovieById.Error != nil {
		return nil, checkMovieById.Error
	}

	deleteMovie := db.Debug().Delete(&movie)
	if deleteMovie.Error != nil {
		return nil, deleteMovie.Error
	}

	return &movie, nil
}
