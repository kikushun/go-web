package controller

import (
	"errors"
	"net/http"

	"github.com/kikuchi/go-web/service"
)

// SaveUser 登録・更新
func SaveUser(w *http.ResponseWriter, r *http.Request) interface{} {

	id := r.PostFormValue("id")
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	resp, err := service.SaveUser(id, name, email, password)
	if err != nil {
		return err
	}
	return resp
}

// SearchUser 検索
func SearchUser(w *http.ResponseWriter, r *http.Request) interface{} {
	ids, ok := r.URL.Query()["id"]
	if !ok {
		return errors.New("パラーメータが見つからない")
	}

	searchResp, err := service.SearchUsers(ids)
	if err != nil {
		return err
	}

	return searchResp
}

// DeleteUser 削除
func DeleteUser(w *http.ResponseWriter, r *http.Request) interface{} {
	IDs, ok := r.URL.Query()["id"]
	if !ok {
		return errors.New("パラーメータが見つからない")
	}

	resp, err := service.DeleteUser(IDs[0])
	if err != nil {
		return err
	}

	return resp
}
