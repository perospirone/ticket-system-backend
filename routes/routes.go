package routes

import (
	"net/http"

	"ticket-system/controllers"
	"ticket-system/middlewares"

	"github.com/gorilla/mux"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func Routes(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/ping", middlewares.JWT(http.HandlerFunc(pong)))
}
