package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"ticket-system/database"
)

func OpenCall(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
	call := &database.Call{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &call)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	result := db.Create(&call)

	if result.Error != nil {
		log.Println(result.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: result.Error.Error()})
		return
	}

	log.Println(call)

	w.WriteHeader(http.StatusCreated)
}

func ListCalls(w http.ResponseWriter, r *http.Request) {
  calls := &[]database.Call{}

  result := db.Find(calls)

	if result.Error != nil {
		log.Println(result.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: result.Error.Error()})
		return
	}

	log.Println(calls)

  json.NewEncoder(w).Encode(calls)
}
