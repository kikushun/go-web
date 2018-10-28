package model

import (
	"encoding/json"
	"net/http"
)

// Handler handler
type Handler struct {
	Main func(*http.Request) interface{}
}

// 共通処理はここ
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// ================
	// 共通処理(前処理)
	// ================

	data := h.Main(r)

	// ================
	// 共通処理(後処理)
	// ================

	// mainProcessの戻り値をResponse.Dataに格納してjsonで返す
	var byteArr []byte
	switch val := data.(type) {
	case error:
		byteArr, _ = json.Marshal(&Response{Status: 9, Msg: val.Error()})
	default:
		byteArr, _ = json.Marshal(&Response{Status: 0, Data: data})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}
