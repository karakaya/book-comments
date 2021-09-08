package model

type Author struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	About string `json:"about"`
	//Birthdate time.Time `json:"birthdate"`
	//Death     time.Time `json:"death"`
}
