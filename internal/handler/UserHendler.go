package handler

import (
	us "BackEnd/internal/userservice"
	"BackEnd/internal/web/users"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StrictUserHendler struct {
	service us.UserService1
}

func NewStrictUserHandler(u us.UserService1) *StrictUserHendler {
	return &StrictUserHendler{service: u}
}

func (u *StrictUserHendler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	user, err := u.service.GetAllUser()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response := users.GetUsers200JSONResponse{}

	for _, val := range user {
		var id int = int(val.Id)
		var email string = val.Email
		var password string = val.Password
		response = append(response, users.User{
			Id:       &id,
			Email:    email,
			Password: password,
		})
	}
	return response, nil
}

func (u *StrictUserHendler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	body := request.Body

	user := new(users.User)

	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Password != "" {
		user.Password = body.Password
	}

	us, err := u.service.CreateUser(user.Email, user.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	apiUser := users.User{
		Email:    us.Email,
		Password: us.Password, // или уберите, если не нужно возвращать пароль
	}

	return users.PostUsers201JSONResponse(apiUser), nil
}

func (u *StrictUserHendler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := int(request.Id)

	if err := u.service.DeleteUser(id); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return users.DeleteUsersId204Response{}, nil
}

func (u *StrictUserHendler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := int(request.Id)

	body := request.Body

	var user us.User
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Password != "" {
		user.Password = body.Password
	}

	if _, err := u.service.UpdataUser(id, user); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	apiUser := users.User{
		Email:    user.Email,
		Password: user.Password, // или уберите, если не нужно отдавать пароль,
	}
	return users.PatchUsersId200JSONResponse(apiUser), nil

}
