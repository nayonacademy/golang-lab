package account

import "context"

type User struct {
	ID string `json:"id"`
	Email string	`json:"email"`
	Password string	`json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	Getuser(ctx context.Context, id string) (string, error)
}
