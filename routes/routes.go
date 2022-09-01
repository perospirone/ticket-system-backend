package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ticket-system/controllers"

	"github.com/gorilla/mux"
)

type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// Articles ...
var articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	articles = []Article{
		Article{Title: "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		Article{Title: "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		Article{Title: "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}
	//fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")

	json.NewEncoder(w).Encode(articles)
}

func Routes(router *mux.Router) {
	router.HandleFunc("/", homePage)
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
