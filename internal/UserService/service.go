package userservice

type UserService interface {
	CreateUser(expression, pass string) (User, error)
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(user string) error
	DeleteUser(id int) error
}

type CUsersServive struct {
	repo UsersRepository
}

func NewUserService(r UsersRepository) UserService {
	return &CUsersServive{repo: r}

}

// CreateUser implements UserService.
func (c *CUsersServive) CreateUser(expression string, pass string) (User, error) {
	ur := User{Email: expression, Password: pass}
	if err := c.repo.CreateUser(ur); err != nil {
		return User{}, err
	}
	return ur, nil
}

// DeleteUser implements UserService.
func (c *CUsersServive) DeleteUser(id int) error {
	return c.repo.DeleteUser(id)
}

// GetAllUser implements UserService.
func (c *CUsersServive) GetAllUser() ([]User, error) {
	return c.repo.GetAllUser()
}

// GetUser implements UserService.
func (c *CUsersServive) GetUser(id int) (User, error) {
	return c.repo.GetUser(id)
}

// UpdataUser implements UserService.
func (c *CUsersServive) UpdataUser(user string) error {
	panic("unimplemented")
}
