package main

import (
	"log"
	"net/http"
	"ticket-system/database"
	"ticket-system/routes"

	"github.com/gorilla/mux"
)

func main() {
	db := database.Connection()
	defer db.Close()

	database.Migrate(db)

	router := mux.NewRouter().StrictSlash(true)

  router.Use(mux.CORSMethodMiddleware(router))

	routes.Routes(router)

  log.Println("initializing server in port 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
