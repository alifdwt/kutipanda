package models

type Artist struct {
	ID        int      `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name" gorm:"size:50;not null"`
	Slug      string   `json:"slug" gorm:"unique;not null"`
	ImageURL  string   `json:"image_url" gorm:"not null"`
	Sex       string   `json:"sex" gorm:"not null"`
	CountryID int      `json:"country_id" gorm:"not null"`
	Country   *Country `json:"country" gorm:"foreignkey:CountryID"`
	Songs     []*Song  `json:"songs" gorm:"many2many:artist_songs;"`
}
