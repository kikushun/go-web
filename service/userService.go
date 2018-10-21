package service

import (
	"net/http"
	"net/url"

	"github.com/kikuchi/go-web/model"
)

// SaveUser ユーザ登録・更新
func SaveUser(user *model.User) (*model.Response, error) {

	escapeURL := "http://localhost:9200/test/doc/" + url.PathEscape(user.Email)

	resp := &model.Response{}
	if err := HTTPRequest(http.MethodPut, escapeURL, user, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
