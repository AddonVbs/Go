package userservice

import "time"

type User struct {
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     string     `json:"email"`
	Id        *int       `json:"id,omitempty"`
	Password  string     `json:"password"`
	UpdatedAt *time.Time `json:"updated_at"`
}
