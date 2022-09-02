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
	router.Handle("/call", middlewares.JWT(http.HandlerFunc(controllers.OpenCall))).Methods("POST")
	router.Handle("/call/list", middlewares.JWT(http.HandlerFunc(controllers.ListCalls))).Methods("GET")
}
