package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kikuchi/go-web/model"
)

// Handler handler
type Handler struct {
	mainProcess func(http.ResponseWriter, *http.Request)
}

// 共通処理はここ
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 共通処理(前処理)

	h.mainProcess(w, r)

	// 共通処理(後処理)

}

// ConvertRequestBodyToModel リクエストボディを引数のモデルに変換
func ConvertRequestBodyToModel(reqBody io.ReadCloser, model interface{}) error {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, model)
}

// ReturnByJSON jsonでレスポンスを返す
func ReturnByJSON(w http.ResponseWriter, result interface{}) {

	respModel := &model.Response{Status: 0, Result: result}
	byteArr, err := json.Marshal(respModel)
	if err != nil {
		respModel.Status = 9
		respModel.Msg = err.Error()
		respModel.Result = ""
		byteArr, _ = json.Marshal(respModel)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}

// ReturnError エラー時のレスポンス
func ReturnError(w http.ResponseWriter, errMsg string) {
	respModel := &model.Response{Status: 9, Msg: errMsg}
	byteArr, _ := json.Marshal(respModel)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}
