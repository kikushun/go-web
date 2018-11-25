package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kikuchi/go-web/config"

	"golang.org/x/crypto/bcrypt"

	"github.com/kikuchi/go-web/model"
)

// SaveUser ユーザ登録・更新
func SaveUser(ID, name, email, password string) (*model.ElasResp, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{ID: ID, Name: name, Email: email, Password: hash}

	URL := fmt.Sprintf("%s/%s/%s/%s", config.BaseURL, config.UserIndex, config.UserType, url.PathEscape(ID))

	resp := &model.ElasResp{}
	if err := HTTPRequest(http.MethodPut, URL, user, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUserByIDAndPassword ログイン認証時、IDとパスワードでユーザ検索
func GetUserByIDAndPassword(id, password string) (*model.User, error) {

	searchReq := &model.SearchReq{}
	searchReq.AddMust("term", "_id", id)
	URL := fmt.Sprintf("%s/%s/_search", config.BaseURL, config.UserIndex)
	searchResp := &model.SearchResp{}
	if err := HTTPRequest(http.MethodPost, URL, searchReq, searchResp); err != nil {
		return nil, err
	}

	if len(searchResp.Hits.Hits) != 1 {
		return nil, errors.New("user取得失敗")
	}

	user := &searchResp.Hits.Hits[0].User
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

// SearchUsers 引数のIDのデータを取得
func SearchUsers(IDs []string) (*model.SearchResp, error) {

	searchReq := &model.SearchReq{}
	for _, id := range IDs {
		searchReq.AddShould("term", "_id", id)
	}

	URL := fmt.Sprintf("%s/%s/_search", config.BaseURL, config.UserIndex)

	searchResp := &model.SearchResp{}
	if err := HTTPRequest(http.MethodPost, URL, searchReq, searchResp); err != nil {
		return nil, err
	}
	return searchResp, nil
}

// DeleteUser ユーザ削除
func DeleteUser(ID string) (*model.ElasResp, error) {

	URL := fmt.Sprintf("%s/%s/%s/%s", config.BaseURL, config.UserIndex, config.UserType, url.PathEscape(ID))

	resp := &model.ElasResp{}
	if err := HTTPRequest(http.MethodDelete, URL, nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
