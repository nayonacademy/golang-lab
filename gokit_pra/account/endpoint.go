package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
}

func MakeEndpoints (s Service) Endpoints{
	return Endpoints{
		CreateUser: makeCreateUserEndpoints(s),
		GetUser:    makeGetUserEndpoints(s),
	}
}

func makeCreateUserEndpoints(s Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{})(interface{}, error){
		req := request.(CreateUserRequest)
		ok , err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok:ok}, err
	}
}

func makeGetUserEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.ID)
		return GetUserResponse{Email:email}, err
	}
}