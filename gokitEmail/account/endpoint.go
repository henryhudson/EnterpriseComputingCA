package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	// contains methods that are exposed to the public
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password, req.Body, req.ToEmail)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)
		body, err := s.GetUser(ctx, req.Id)
		toEmail, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email:   email,
			Body:    body,
			ToEmail: toEmail,
		}, err
	}
}
