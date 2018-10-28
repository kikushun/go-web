package service

import (
	"net/http"
	"net/url"

	"github.com/kikuchi/go-web/model"
)

// SaveUsers ユーザ登録・更新
func SaveUsers(user *model.User) (*model.SaveResp, error) {

	escapeURL := "http://localhost:9200/test/doc/" + url.PathEscape(user.Email)

	resp := &model.SaveResp{}
	if err := HTTPRequest(http.MethodPut, escapeURL, user, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// SearchUsers 引数のIDのデータを取得
func SearchUsers(IDs []string) (*model.SearchResp, error) {

	searchReq := &model.SearchReq{}
	for _, id := range IDs {
		searchReq.AddShould("term", "_id", id)
	}

	searchResp := &model.SearchResp{}
	const URL = "http://localhost:9200/test/_search"
	if err := HTTPRequest(http.MethodPost, URL, searchReq, searchResp); err != nil {
		return nil, err
	}
	return searchResp, nil
}
