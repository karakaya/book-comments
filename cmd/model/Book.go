package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `json:"title"`
	About      string `json:"about"`
	CategoryID uint   `json:"category"`
	Category   Category
	AuthorID   uint `json:"author"`
	Author     Author
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
