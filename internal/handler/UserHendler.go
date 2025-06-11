package handler

import (
	us "BackEnd/internal/UserService"
	"BackEnd/internal/web/users"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StrictUserHendler struct {
	service us.UserService
}

func NewStrictUserHandler(u us.UserService) *StrictUserHendler {
	return &StrictUserHendler{service: u}
}

func (u *StrictUserHendler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	user, err := u.service.GetAllUser()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response := users.GetUsers200JSONResponse{}

	for _, val := range user {
		var id int = int(*val.Id)
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

}

func (u *StrictUserHendler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {

}

func (u *StrictUserHendler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {

}
