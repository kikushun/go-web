package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kikuchi/go-web/model"
	"github.com/kikuchi/go-web/service"
)

// handler handler
type handler struct {
	method   string
	mainFunc func(*http.Request) interface{}
}

// 共通処理はここ
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// ================
	// 共通処理(前処理)
	// ================

	// リクエストメソッドチェック
	if h.method != r.Method {
		byteArr, _ := json.Marshal(&model.Response{Status: 9, Msg: fmt.Sprintf("No handler found for uri [%s] and method [%s]", r.RequestURI, r.Method)})
		w.Write(byteArr)
		return
	}

	// 認証チェック
	if err := service.Authenticate(r); err != nil {
		byteArr, _ := json.Marshal(&model.Response{Status: 9, Msg: err.Error()})
		w.Write(byteArr)
		return
	}

	// ================
	// メイン処理
	// ================
	data := h.mainFunc(r)

	// ================
	// 共通処理(後処理)
	// ================

	// mainProcessの戻り値をResponse.Dataに格納してjsonで返す
	w.Header().Set("Content-Type", "application/json")
	var byteArr []byte
	switch val := data.(type) {
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		byteArr, _ = json.Marshal(&model.Response{Status: 9, Msg: val.Error()})
	default:
		byteArr, _ = json.Marshal(&model.Response{Status: 0, Data: data})
	}

	w.Write(byteArr)
}

// h Handlerを返す
func h(method string, mainFunc func(*http.Request) interface{}) *handler {

	return &handler{
		method:   method,
		mainFunc: mainFunc,
	}
}
