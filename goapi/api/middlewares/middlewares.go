package middlewares

import "net/http"

func SetMiddlewaresJSON(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		next(w, r)
	}
}