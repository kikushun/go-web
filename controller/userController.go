package controller

import (
	"fmt"
	"net/http"
)

// UserController aa
func UserController(mux *http.ServeMux) {
	mux.Handle("/user", Handler{getUser})
}

// /user
func getUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello!!!!")
}
