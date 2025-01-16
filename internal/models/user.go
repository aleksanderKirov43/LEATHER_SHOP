package models

// Структура для хранения полей с информацией о пользователе
type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Wishlist  int    `json:"wishlist"`
	Cart      int    `json:"cart"`
}
