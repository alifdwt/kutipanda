package models

import "time"

type Song struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title" gorm:"unique;not null"`
	Lyrics        string    `json:"lyrics" gorm:"not null"`
	Year          int       `json:"year" gorm:"not null"`
	AlbumImageUrl string    `json:"album_image_url" gorm:"not null"`
	Language      string    `json:"language" gorm:"not null"`
	UserID        int       `json:"user_id" gorm:"not null"`
	User          User      `json:"user"`
	Artists       []*Artist `json:"artists" gorm:"many2many:artist_songs;"`
	Movie         []*Movie  `json:"movies" gorm:"many2many:movie_songs;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
