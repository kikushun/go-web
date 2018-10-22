package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ConvertToStruct 引数で渡された構造体に変換する
func ConvertToStruct(reader io.Reader, model interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, model)
}
