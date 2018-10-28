package controller

import (
	"fmt"
	"net/http"

	"github.com/kikuchi/go-web/model"
)

// CategoryController カテゴリコントローラ
func CategoryController(mux *http.ServeMux) {

	mux.Handle("/category", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		values := req.URL.Query()
		ids, ok := values["id"]
		if !ok {
			return nil, nil
		}
		for _, val := range ids {
			fmt.Println(val)
		}

		searchReq := model.Search{}
		searchReq.AddMust("term", "id", "aaaaaaa")

		fmt.Println(searchReq.Musts[0].Term["id"])

		return nil, nil
	}})

	mux.Handle("/categories", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		return nil, nil
	}})
}
