package controller

import (
	"net/http"

	"github.com/kikuchi/go-web/model"
)

// CategoryController カテゴリコントローラ
func CategoryController(mux *http.ServeMux) {

	mux.Handle("/category", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		return nil, nil
	}})

	mux.Handle("/categories", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		return nil, nil
	}})
}
