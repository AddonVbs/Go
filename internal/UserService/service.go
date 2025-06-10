package userservice

type UserService interface {
	CreateUser(expression, password string) error
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(user string) error
	UpdataPass(Pass string) error
	DeleteUser(id int) error
}

type UsersServiveDb struct {
	repo UsersRepository
}

func NewUserService() {

}
