package models

import "time"

type Quote struct {
	ID                   int       `json:"id" gorm:"primaryKey"`
	QuoteText            string    `json:"quote_text" gorm:"not null"`
	QuoteTransliteration string    `json:"quote_transliteration"`
	LanguageID           int       `json:"language_id" gorm:"not null"`
	Language             *Language `json:"language" gorm:"foreignkey:LanguageID"`
	MovieID              int       `json:"movie_id" gorm:"not null"`
	Movie                *Movie    `json:"movie" gorm:"foreignkey:MovieID"`
	UserID               int       `json:"user_id" gorm:"not null"`
	User                 User      `json:"user"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
