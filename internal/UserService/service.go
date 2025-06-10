package userservice

import "gorm.io/gorm"

type UserService interface {
	CreateUser(expression, password string) error
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(user string) error
	UpdataPass(Pass string)
	DeleteUser(id int) error
}

type UsersServiveDb struct {
	db *gorm.DB
}

func NewUserService()
