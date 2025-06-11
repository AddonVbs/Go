package handler

import (
	us "BackEnd/internal/UserService"
	"BackEnd/internal/web/users"
	"context"
)

type StrictUserHendler struct {
	service us.UserService
}

func NewStrictUserHandler(u us.UserService) *StrictUserHendler {
	return &StrictUserHendler{service: u}
}

func (u *StrictUserHendler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {

}

func (u *StrictUserHendler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

}

func (u *StrictUserHendler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {

}

func (u *StrictUserHendler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {

}
