package model

import (
	"time"
)

type Author struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Birthdate time.Time
	Death     time.Time
	Books     []Book
}
