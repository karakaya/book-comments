package model

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `json:"user"`
	User    User
	BookID  uint `json:"book"`
	Book    Book
	Comment string `json:"comment"`
}
