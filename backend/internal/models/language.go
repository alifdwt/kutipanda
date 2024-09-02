package models

type Language struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}
