package main

import (
	"net/http"

	"github.com/kikuchi/go-web/controller"
)

func main() {
	mux := http.NewServeMux()
	controller.UserController(mux)
	controller.CategoryController(mux)

	http.ListenAndServe(":9090", mux)
}
