package controller

import (
	"errors"
	"net/http"

	"github.com/kikuchi/go-web/service"
)

// SaveUser 登録・更新
func SaveUser(req *http.Request) interface{} {

	id := req.PostFormValue("id")
	name := req.PostFormValue("name")
	email := req.PostFormValue("email")
	password := req.PostFormValue("password")

	resp, err := service.SaveUser(id, name, email, password)
	if err != nil {
		return err
	}
	return resp
}

// SearchUser 検索
func SearchUser(req *http.Request) interface{} {
	ids, ok := req.URL.Query()["id"]
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
func DeleteUser(req *http.Request) interface{} {
	IDs, ok := req.URL.Query()["id"]
	if !ok {
		return errors.New("パラーメータが見つからない")
	}

	resp, err := service.DeleteUser(IDs[0])
	if err != nil {
		return err
	}

	return resp
}
