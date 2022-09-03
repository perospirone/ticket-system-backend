package main

import (
	"log"
	"net/http"
	"ticket-system/database"
	"ticket-system/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db := database.Connection()
	defer db.Close()

	database.Migrate(db)

	router := mux.NewRouter().StrictSlash(true)

	router.Use(mux.CORSMethodMiddleware(router))

	routes.Routes(router)

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    AllowedHeaders: []string{"*"},
  })

	log.Println("initializing server in port 8000")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
}
