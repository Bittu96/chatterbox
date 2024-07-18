package user

type User struct {
	UserId    int64       `sql:"user_id" json:"user_id"`
	Username  string      `json:"username" validate:"required"`
	Email     string      `json:"email" validate:"required"`
	Password  string      `json:"password" validate:"required"`
	Role      string      `json:"role" validate:"required"`
	CreatedAt interface{} `sql:"created_at" json:"created_at"`
	UpdatedAt interface{} `sql:"updated_at" json:"updated_at"`
}

// func (u User) Check() {

// }

// func (u User) Add() {

// }

// func (u User) Delete() {

// }

// func (u User) Update() {

// }
