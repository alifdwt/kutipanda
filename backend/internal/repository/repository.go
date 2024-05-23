package repository

import "gorm.io/gorm"

type Repositories struct {
	User  UserRepository
	Movie MovieRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:  NewUserRepository(db),
		Movie: NewMovieRepository(db),
	}
}
