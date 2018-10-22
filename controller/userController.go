package controller

import (
	"net/http"

	"github.com/kikuchi/go-web/util"

	"github.com/kikuchi/go-web/service"

	"github.com/kikuchi/go-web/model"
)

// UserController ユーザコントローラ
func UserController(mux *http.ServeMux) {

	// =================
	// ユーザ登録・更新
	// =================
	mux.Handle("/user/save", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		user := &model.User{}
		if err := util.ConvertToStruct(req.Body, user); err != nil {
			return nil, err
		}
		return service.SaveUser(user)
	}})

	// =================
	// ユーザ情報取得
	// =================
	mux.Handle("/user", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		return nil, nil
	}})

	// ================
	// ユーザリスト取得
	// ================
	mux.Handle("/users", model.Handler{MainProcess: func(req *http.Request) (interface{}, error) {
		return nil, nil
	}})
}
