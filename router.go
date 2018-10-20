package main

import (
	"net/http"

	"github.com/kikuchi/go-web/controller"
)

// ServeMux aaaa
func ServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	controller.UserController(mux)
	controller.CategoryController(mux)
	return mux
}
