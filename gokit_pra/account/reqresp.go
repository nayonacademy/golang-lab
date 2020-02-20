package account

import (
	"context"
	"encoding/json"
	"net/http"
)

type(
	CreateUserRequest struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		ok string `json:"ok"`
	}
	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeRespose(ctx context.Context, w http.ResponseWriter, response interface{}){
	return json.NewEncoder(w).Encode(response)
}

func decodeLUserReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req CreateUserRequest
	err := json.NewEncoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
}

func deoceEmailReq(ctx context.Context, r *http.Request)(interface{}, error)  {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{ID:vars["id"],}
	return req, nil
}