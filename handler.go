package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kikuchi/go-web/model"
)

// handler handler
type handler struct {
	method   string
	mainFunc func(*http.Request) interface{}
}

// 共通処理はここ
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// ================
	// 共通処理(前処理)
	// ================
	if h.method != r.Method {
		http.Error(w, fmt.Sprintf("No handler found for uri [%s] and method [%s]", r.RequestURI, r.Method), 400)
		return
	}

	data := h.mainFunc(r)

	// ================
	// 共通処理(後処理)
	// ================

	// mainProcessの戻り値をResponse.Dataに格納してjsonで返す
	var byteArr []byte
	switch val := data.(type) {
	case error:
		byteArr, _ = json.Marshal(&model.Response{Status: 9, Msg: val.Error()})
	default:
		byteArr, _ = json.Marshal(&model.Response{Status: 0, Data: data})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}

// h Handlerを返す
func h(method string, mainFunc func(*http.Request) interface{}) *handler {

	return &handler{
		method:   method,
		mainFunc: mainFunc,
	}
}
