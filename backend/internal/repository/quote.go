package repository

import (
	"github.com/alifdwt/kutipanda-backend/internal/domain/requests/quote"
	"github.com/alifdwt/kutipanda-backend/internal/models"
	"gorm.io/gorm"
)

type quoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *quoteRepository {
	return &quoteRepository{db: db}
}

func (r *quoteRepository) CreateQuote(userId int, quote quote.CreateQuoteRequest) (*models.Quote, error) {
	var quoteModel models.Quote

	db := r.db.Model(&quoteModel)

	quoteModel.QuoteText = quote.QuoteText
	quoteModel.QuoteTransliteration = quote.QuoteTransliteration
	quoteModel.MovieID = quote.MovieID
	quoteModel.LanguageID = quote.LanguageID
	quoteModel.UserID = userId

	if err := db.Debug().Create(&quoteModel).Error; err != nil {
		return nil, err
	}

	return &quoteModel, nil
}

func (r *quoteRepository) GetQuoteAll() (*[]models.Quote, error) {
	var quotes []models.Quote

	db := r.db.Model(&quotes)

	if err := db.Debug().Preload("Movie.Country").Preload("Language").Preload("User").Find(&quotes).Error; err != nil {
		return nil, err
	}

	return &quotes, nil
}

func (r *quoteRepository) GetRandomQuote(count int) (*[]models.Quote, error) {
	var quotes []models.Quote

	db := r.db.Model(&quotes)

	result := db.Debug().Preload("Movie.Country").Preload("Language").Preload("User").Order("RANDOM()").Limit(count).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quotes, nil
}

func (r *quoteRepository) DeleteQuoteById(id int) (*models.Quote, error) {
	var quote models.Quote

	db := r.db.Model(quote)

	checkQuoteById := db.Debug().Where("id = ?", id).First(&quote)
	if checkQuoteById.Error != nil {
		return nil, checkQuoteById.Error
	}

	deleteQuote := db.Debug().Delete(&quote)
	if deleteQuote.Error != nil {
		return nil, deleteQuote.Error
	}

	return &quote, nil
}
