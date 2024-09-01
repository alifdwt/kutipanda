package repository

import "gorm.io/gorm"

type Repositories struct {
	User            UserRepository
	Movie           MovieRepository
	Song            SongRepository
	SongTranslation SongTranslationRepository
	Country         CountryRepository
	Quote           QuoteRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:            NewUserRepository(db),
		Movie:           NewMovieRepository(db),
		Song:            NewSongRepository(db),
		SongTranslation: NewSongTranslationRepository(db),
		Country:         NewCountryRepository(db),
		Quote:           NewQuoteRepository(db),
	}
}
