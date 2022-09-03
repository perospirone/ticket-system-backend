package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"ticket-system/database"
)

var db = database.Connection()

func Register(w http.ResponseWriter, r *http.Request) {
	user := &database.User{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	result := db.Create(user)

	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: result.Error.Error()})
		return
	}

	token, err := createTokenJWT(user.Name, user.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	response := responseToken{token}

	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
  //w.Header().Set("Content-Type", "application/json")

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")

	body := &database.User{}
	user := &database.User{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	result := db.Take(&user, "email = ?", body.Email)

	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: result.Error.Error()})
		return
	}

	if body.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responseError{Error: "wrong password"})
		return
	}

	token, err := createTokenJWT(user.Name, user.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	response := responseToken{token}

	json.NewEncoder(w).Encode(response)
}
