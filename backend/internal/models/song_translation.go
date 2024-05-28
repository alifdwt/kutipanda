package models

import "time"

type SongTranslation struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	TranslatedTitle  string    `json:"translated_title" gorm:"unique;not null"`
	Slug             string    `json:"slug" gorm:"unique;not null"`
	TranslatedLyrics string    `json:"translated_lyrics" gorm:"unique;not null"`
	Language         string    `json:"language" gorm:"not null"`
	UserID           int       `json:"user_id" gorm:"not null"`
	User             User      `json:"user"`
	SongID           int       `json:"song_id" gorm:"not null"`
	Song             Song      `json:"song"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
