package account

import (
	"context"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints)(http.Handler){
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeLUserReq,
		encodeRespose,
		))
	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		encodeRespose,
		 ))
	return r
}


func commonMiddleware(next http.Handler) http.Handler{
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type","application/json")
		next.ServeHTTP(w, r)
	})
}