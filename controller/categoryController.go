package controller

import (
	"fmt"
	"net/http"
)

// CategoryController aa
func CategoryController(mux *http.ServeMux) {
	mux.Handle("/category", Handler{getCategory})
	mux.Handle("/categories", Handler{getCategories})
}

// /user
func getCategory(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "category")
}

func getCategories(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "categories")
}
