package model

import (
	"encoding/json"
	"net/http"
)

// Handler handler
type Handler struct {
	MainProcess func(*http.Request) (interface{}, error)
}

// 共通処理はここ
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// ================
	// 共通処理(前処理)
	// ================

	result, err := h.MainProcess(r)

	// ================
	// 共通処理(後処理)
	// ================

	// mainProcessの戻り値の構造体をResponse.Resultに格納してjsonで返す
	respModel := &Response{Status: 0, Result: result}
	if err != nil {
		respModel = &Response{Status: 9, Msg: err.Error()}
	}
	byteArr, _ := json.Marshal(respModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}
