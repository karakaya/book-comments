package model

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}
