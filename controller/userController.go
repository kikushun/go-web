package controller

import (
	"errors"
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
	mux.Handle("/user/save", model.Handler{Main: func(req *http.Request) interface{} {
		user := &model.User{}
		if err := util.ConvertToStruct(req.Body, user); err != nil {
			return err
		}
		model, err := service.SaveUsers(user)
		if err != nil {
			return err
		}
		return model
	}})

	// =================
	// ユーザ情報取
	// =================
	mux.Handle("/user/search", model.Handler{Main: func(req *http.Request) interface{} {

		ids, ok := req.URL.Query()["id"]
		if !ok {
			return errors.New("パラーメータが見つからない")
		}

		searchResp, err := service.SearchUsers(ids)
		if err != nil {
			return err
		}

		return searchResp
	}})

	// =================
	// ユーザ削除
	// =================
	mux.Handle("/user/delete", model.Handler{Main: func(req *http.Request) interface{} {
		return nil
	}})
}
