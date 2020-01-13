package main

import "github.com/gorilla/mux"

import "net/http"

import "fmt"

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeFunction)
	// r.HandleFunc("/products", ArticleCategoryHandle)
	// r.HandleFunc("/product/{category}/", ArticleCategoryHandle)
	// r.HandleFunc("/category/{category}/{id:[0-9]+}", ArticleCategoryHandle)
	// http.Handle("/", r)

	// matching router
	r := mux.NewRouter()
	r.Host("www.example.com")
	r.Host("{subdomain:[a-z]+}.example.com")
	r.PathPrefix("/product")
	r.Methods("GET", "POST")
	r.Schemes("https")
	r.Headers("X-Requested-With", "XMLHttpRequest")
	r.Queries("key", "value")
	r.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.ProtoMajor == 0
	})
	r.HandleFunc("/hello", ArticleCategoryHandle).
		Host("www.example.com").
		Methods("GET").
		Schemes("http")
	r.PathPrefix("/").Handler(ArticleCategoryHandle)
	// sub router
	s := r.Host("www.example.com").Subrouter()
	s.HandleFunc("/one", ArticleCategoryHandle)
}

func HomeFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello home page")
}

func ArticleCategoryHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
