package model

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
}
