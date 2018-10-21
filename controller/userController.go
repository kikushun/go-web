package controller

import (
	"net/http"

	"github.com/kikuchi/go-web/service"

	"github.com/kikuchi/go-web/model"
)

// UserController ユーザコントローラ
func UserController(mux *http.ServeMux) {

	// =================
	// ユーザ登録・更新
	// =================
	mux.Handle("/user/save", Handler{func(w http.ResponseWriter, req *http.Request) {

		user := &model.User{}
		if err := ConvertRequestBodyToModel(req.Body, user); err != nil {
			ReturnError(w, err.Error())
		}

		resp, err := service.SaveUser(user)
		if err != nil {
			ReturnError(w, err.Error())
		}

		ReturnByJSON(w, resp)
	}})

	// =================
	// ユーザ情報取得
	// =================
	mux.Handle("/user", Handler{func(w http.ResponseWriter, req *http.Request) {
		return
	}})

	// ================
	// ユーザリスト取得
	// ================
	mux.Handle("/users", Handler{func(w http.ResponseWriter, req *http.Request) {
		return
	}})
}
