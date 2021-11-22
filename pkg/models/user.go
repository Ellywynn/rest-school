package models

type User struct {
	Id       uint   `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"hashed_password" binding:"required"`
}
