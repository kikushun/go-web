package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kikuchi/go-web/model"
	"github.com/kikuchi/go-web/service"
)

// Login login
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// リクエストメソッドチェック
	if http.MethodPost != r.Method {
		byteArr, _ := json.Marshal(&model.Response{Status: 9, Msg: fmt.Sprintf("No handler found for uri [%s] and method [%s]", r.RequestURI, r.Method)})
		w.Write(byteArr)
		return
	}

	// elasのユーザに存在するかチェック
	id := r.PostFormValue("id")
	password := r.PostFormValue("password")
	user, err := service.GetUserByIDAndPassword(id, password)
	if err != nil {
		byteArr, _ := json.Marshal(&model.Response{Status: 9, Msg: err.Error()})
		w.Write(byteArr)
		return
	}

	// トークン発行
	token, err := service.CreateToken(user)
	if err != nil {
		byteArr, _ := json.Marshal(&model.Response{Status: 9, Msg: err.Error()})
		w.Write(byteArr)
		return
	}

	// トークンをDataに詰めてレスポンス
	byteArr, _ := json.Marshal(&model.Response{Status: 0, Data: fmt.Sprintf("Bearer %s", token)})
	w.Write(byteArr)
}
