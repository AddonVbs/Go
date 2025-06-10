package userservice

type UserService interface {
	CreateUser(expression, password string) error
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(user string) error
	UpdataPass(Pass string) error
	DeleteUser(id int) error
}

type CUsersServive struct {
	repo UsersRepository
}

// CreateUser implements UserService.
func (c *CUsersServive) CreateUser(expression string, password string) error {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (c *CUsersServive) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetAllUser implements UserService.
func (c *CUsersServive) GetAllUser() ([]User, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (c *CUsersServive) GetUser(id int) (User, error) {
	panic("unimplemented")
}

// UpdataPass implements UserService.
func (c *CUsersServive) UpdataPass(Pass string) error {
	panic("unimplemented")
}

// UpdataUser implements UserService.
func (c *CUsersServive) UpdataUser(user string) error {
	panic("unimplemented")
}

func NewUserService(r UsersRepository) UserService {
	return &CUsersServive{repo: r}

}
