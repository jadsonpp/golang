package controllers

import (
	"net/http"
)

// Register a new controller

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
