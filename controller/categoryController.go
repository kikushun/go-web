package controller

import (
	"net/http"
)

// CategoryController カテゴリコントローラ
func CategoryController(mux *http.ServeMux) {

	mux.Handle("/category", Handler{func(w http.ResponseWriter, req *http.Request) {
		return
	}})

	mux.Handle("/categories", Handler{func(w http.ResponseWriter, req *http.Request) {
		return
	}})
}
