package repository

import (
	"fmt"

	songtranslation "github.com/alifdwt/kutipanda-backend/internal/domain/requests/song_translation"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type songTranslationRepository struct {
	db *gorm.DB
}

func NewSongTranslationRepository(db *gorm.DB) *songTranslationRepository {
	return &songTranslationRepository{db}
}

func (r *songTranslationRepository) CreateSongTranslation(userId int, request songtranslation.CreateSongTranslationRequest) (*models.SongTranslation, error) {
	var songTranslationModel models.SongTranslation

	db := r.db.Model(&songTranslationModel)

	songTranslationModel.TranslatedTitle = request.TranslatedTitle
	songTranslationModel.TranslatedLyrics = request.TranslatedLyrics
	songTranslationModel.Language = request.Language

	songTranslationModel.UserID = userId
	songTranslationModel.SongID = request.SongID
	songTranslationModel.Slug = fmt.Sprintf("%s-%d", slug.Make(request.TranslatedTitle), request.SongID)

	result := db.Debug().Create(&songTranslationModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songTranslationModel, nil
}

func (r *songTranslationRepository) GetSongTranslationAll(limit int, offset int, order string) (*[]models.SongTranslation, error) {
	var songTranslations []models.SongTranslation

	db := r.db.Model(&songTranslations)

	result := db.Debug().
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&songTranslations)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songTranslations, nil
}

func (r *songTranslationRepository) GetSongTranslationById(id int) (*models.SongTranslation, error) {
	var songTranslationModel models.SongTranslation

	db := r.db.Model(&songTranslationModel)

	result := db.Debug().Where("id = ?", id).First(&songTranslationModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songTranslationModel, nil
}

func (r *songTranslationRepository) GetSongTranslationBySlug(slug string) (*models.SongTranslation, error) {
	var songTranslationModel models.SongTranslation

	db := r.db.Model(&songTranslationModel)

	result := db.Debug().Where("slug = ?", slug).First(&songTranslationModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &songTranslationModel, nil
}

func (r *songTranslationRepository) UpdateSongTranslation(userId int, songTranslationId int, updatedSongTranslation songtranslation.UpdateSongTranslationRequest) (*models.SongTranslation, error) {
	var songTranslationModel models.SongTranslation

	db := r.db.Model(&songTranslationModel)

	checkSongTranslationById := db.Debug().Where("id = ? AND user_id = ?", songTranslationId, userId).First(&songTranslationModel)
	if checkSongTranslationById.Error != nil {
		return nil, checkSongTranslationById.Error
	}

	songTranslationModel.TranslatedTitle = updatedSongTranslation.TranslatedTitle
	songTranslationModel.TranslatedLyrics = updatedSongTranslation.TranslatedLyrics
	songTranslationModel.Language = updatedSongTranslation.Language

	songTranslationModel.Slug = fmt.Sprintf("%s-%d", slug.Make(updatedSongTranslation.TranslatedTitle), songTranslationId)

	updateSongTranslation := db.Debug().Updates(&songTranslationModel)
	if updateSongTranslation.Error != nil {
		return nil, updateSongTranslation.Error
	}

	return &songTranslationModel, nil
}

func (r *songTranslationRepository) DeleteSongTranslation(userId int, songTranslationId int) (*models.SongTranslation, error) {
	var songTranslationModel models.SongTranslation

	db := r.db.Model(&songTranslationModel)

	checkSongTranslationById := db.Debug().Where("id = ? AND user_id = ?", songTranslationId, userId).First(&songTranslationModel)
	if checkSongTranslationById.Error != nil {
		return nil, checkSongTranslationById.Error
	}

	deleteSongTranslation := db.Debug().Delete(&songTranslationModel)
	if deleteSongTranslation.Error != nil {
		return nil, deleteSongTranslation.Error
	}

	return &songTranslationModel, nil
}
