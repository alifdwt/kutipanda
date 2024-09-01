package models

import "time"

type Actor struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:50;unique;not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	Sex       string    `json:"sex" gorm:"not null"`
	CountryID int       `json:"country_id" gorm:"not null"`
	Country   *Country  `json:"country" gorm:"foreignkey:CountryID"`
	Movies    []*Movie  `json:"movies" gorm:"many2many:movie_actors;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
