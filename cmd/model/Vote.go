package model

type Vote struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User

	BookID uint
	Book   Book
}
