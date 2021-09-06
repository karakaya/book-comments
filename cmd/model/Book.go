package model

import (
	"gorm.io/gorm"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Pages     uint
	Category  Category
	DeletedAt gorm.DeletedAt
}
