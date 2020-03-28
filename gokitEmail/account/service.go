package account

import "context"

type Service interface {
	CreateUser(ctx context.Context,
		email string,
		password string,
		body string,
		toEmail string) (string, error)
	GetUser(ctx context.Context,
		id string) (string, error)
}
