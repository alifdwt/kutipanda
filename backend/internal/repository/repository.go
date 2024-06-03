package repository

import "gorm.io/gorm"

type Repositories struct {
	User            UserRepository
	Movie           MovieRepository
	Song            SongRepository
	SongTranslation SongTranslationRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:            NewUserRepository(db),
		Movie:           NewMovieRepository(db),
		Song:            NewSongRepository(db),
		SongTranslation: NewSongTranslationRepository(db),
	}
}
