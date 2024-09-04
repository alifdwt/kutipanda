package repository

import (
	"fmt"

	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/artist"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type artistRepository struct {
	db *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *artistRepository {
	return &artistRepository{db}
}

func (r *artistRepository) CreateArtist(request artist.CreateArtistRequest) (*models.Artist, error) {
	var artistModel models.Artist

	dbArtist := r.db.Model(&artistModel)

	artistModel.Name = request.Name
	artistModel.Sex = request.Sex
	artistModel.CountryID = request.CountryID

	artistModel.Slug = fmt.Sprintf("%s-%d", slug.Make(request.Name), request.CountryID)

	if err := dbArtist.Create(&artistModel).Error; err != nil {
		return nil, err
	}

	return &artistModel, nil
}

func (r *artistRepository) GetArtistAll(limit int, offset int, order string) (*[]models.Artist, error) {
	var artistsModel []models.Artist

	dbArtist := r.db.Model(&artistsModel)

	if err := dbArtist.Debug().Preload("Country").Order(order).Limit(limit).Offset(offset).Find(&artistsModel).Error; err != nil {
		return nil, err
	}

	return &artistsModel, nil
}

func (r *artistRepository) GetArtistBySlug(slug string) (*models.Artist, error) {
	var artistModel models.Artist

	dbArtist := r.db.Model(&artistModel)

	if err := dbArtist.Debug().Preload("Country").Preload("Songs").Where("slug = ?", slug).First(&artistModel).Error; err != nil {
		return nil, err
	}

	return &artistModel, nil
}

func (r *artistRepository) FindArtistsBySlug(slug string, limit int, offset int, order string) (*[]models.Artist, error) {
	var artistsModel []models.Artist

	dbArtist := r.db.Model(&artistsModel)

	result := dbArtist.Debug().
		Where("slug ILIKE ?", slug).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&artistsModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &artistsModel, nil
}

func (r *artistRepository) UpdateArtist(id int, request artist.UpdateArtistRequest) (*models.Artist, error) {
	var artistModel models.Artist

	dbArtist := r.db.Model(&artistModel)

	if err := dbArtist.Debug().Where("id = ?", id).First(&artistModel).Error; err != nil {
		return nil, err
	}

	artistModel.Name = request.Name
	artistModel.Sex = request.Sex
	artistModel.CountryID = request.CountryID

	if err := dbArtist.Debug().Updates(&artistModel).Error; err != nil {
		return nil, err
	}

	return &artistModel, nil
}

func (r *artistRepository) DeleteArtistById(id int) (*models.Artist, error) {
	var artistModel models.Artist

	dbArtist := r.db.Model(&artistModel)

	if err := dbArtist.Debug().Where("id = ?", id).First(&artistModel).Error; err != nil {
		return nil, err
	}

	if err := dbArtist.Debug().Delete(&artistModel).Error; err != nil {
		return nil, err
	}

	return &artistModel, nil
}
