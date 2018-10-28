package controller

import (
	"net/http"

	"github.com/kikuchi/go-web/model"
)

// CategoryController カテゴリコントローラ
func CategoryController(mux *http.ServeMux) {

	mux.Handle("/category", model.Handler{Main: func(req *http.Request) interface{} {
		return nil
	}})

	mux.Handle("/categories", model.Handler{Main: func(req *http.Request) interface{} {
		return nil
	}})
}
