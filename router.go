package main

import (
	"net/http"

	"github.com/kikuchi/controller"
)

// ServeMux aaaa
func ServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	controller.UserController(mux)
	return mux
}
