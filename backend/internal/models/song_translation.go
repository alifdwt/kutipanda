package models

import "time"

type SongTranslation struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	TranslatedTitle  string    `json:"translated_title" gorm:"unique;not null"`
	Slug             string    `json:"slug" gorm:"unique;not null"`
	TranslatedLyrics string    `json:"translated_lyrics" gorm:"unique;not null"`
	LanguageID       int       `json:"language_id" gorm:"not null"`
	Language         Language  `json:"language"`
	UserID           int       `json:"user_id" gorm:"not null"`
	User             User      `json:"user"`
	SongID           int       `json:"song_id" gorm:"not null"`
	Song             Song      `json:"song"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
