package model

import "time"

type Ask struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user"`
	User      User
	CreatedAt time.Time
}
