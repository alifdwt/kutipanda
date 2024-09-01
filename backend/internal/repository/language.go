package repository

import (
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"gorm.io/gorm"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *languageRepository {
	return &languageRepository{db}
}

func (r *languageRepository) CreateLanguage(language models.Language) (*models.Language, error) {
	var languageModel models.Language

	db := r.db.Model(&languageModel)

	if err := db.Debug().Create(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}

func (r *languageRepository) GetLanguageAll() (*[]models.Language, error) {
	var languages []models.Language

	db := r.db.Model(&languages)

	if err := db.Debug().Find(&languages).Error; err != nil {
		return nil, err
	}

	return &languages, nil
}

func (r *languageRepository) GetLanguageByCode(code string) (*models.Language, error) {
	var language models.Language

	db := r.db.Model(&language)

	if err := db.Debug().Where("code = ?", code).First(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}

func (r *languageRepository) DeleteLanguageById(id int) (*models.Language, error) {
	var language models.Language

	db := r.db.Model(&language)

	if err := db.Debug().Where("id = ?", id).Delete(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}
