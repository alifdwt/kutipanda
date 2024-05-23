package models

import "time"

type Movie struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Title          string    `json:"title" gorm:"unique;not null"`
	Slug           string    `json:"slug" gorm:"unique;not null"`
	Description    string    `json:"description" gorm:"size:150;not null"`
	Year           int       `json:"year" gorm:"not null"`
	PosterImageUrl string    `json:"poster_image_url" gorm:"not null"`
	Origin         string    `json:"origin" gorm:"not null"`
	Actors         []*Actor  `json:"actors" gorm:"many2many:movie_actors;"`
	Songs          []*Song   `json:"songs" gorm:"many2many:movie_songs;"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
