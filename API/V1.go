package api

import (
	"net/http"
	user "spotless/controller/User"

	"github.com/gorilla/mux"
)

type RESTApiV1 struct {
	router *mux.Router
}

func (api *RESTApiV1) Server(addr string) error{
	return http.ListenAndServe(addr, api.router)
}

func NewRestApi() *RESTApiV1{
	router := mux.NewRouter()
	api := &RESTApiV1{
		router,
	}

	router.HandleFunc("/user", user.Get).Methods("GET")
	router.HandleFunc("/user", user.Post).Methods("POST")
	router.HandleFunc("/user/{id}", user.GetUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", user.Update).Methods("PUT")
	router.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")

	return api
}