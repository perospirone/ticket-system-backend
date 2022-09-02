package controllers

type responseToken struct {
	Token string `json:"token"`
}

type responseError struct {
	Error string `json:"error"`
}
