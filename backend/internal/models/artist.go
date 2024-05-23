package models

type Artist struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"size:50;unique;not null"`
	ImageURL string  `json:"image_url" gorm:"not null"`
	Sex      string  `json:"sex" gorm:"not null"`
	Origin   string  `json:"origin" gorm:"not null"`
	Songs    []*Song `json:"songs" gorm:"many2many:artist_songs;"`
}
