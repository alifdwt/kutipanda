package models

import "time"

type Song struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title" gorm:"unique;not null"`
	Slug          string    `json:"slug" gorm:"unique;not null"`
	Lyrics        string    `json:"lyrics" gorm:"not null"`
	ReleaseDate   time.Time `json:"release_date" gorm:"not null"`
	AlbumImageUrl string    `json:"album_image_url" gorm:"not null"`
	UserID        int       `json:"user_id" gorm:"not null"`
	User          User      `json:"user"`
	CountryID     int       `json:"country_id" gorm:"not null"`
	Country       *Country  `json:"country" gorm:"foreignkey:CountryID"`
	Artists       []*Artist `json:"artists" gorm:"many2many:artist_songs;"`
	Movie         []*Movie  `json:"movies" gorm:"many2many:movie_songs;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
