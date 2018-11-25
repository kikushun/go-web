package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	// BaseURL elasのURL
	BaseURL string
	// UserIndex userインデックス
	UserIndex string
	// UserType userタイプ
	UserType string
	// SecretKey 認証キー
	SecretKey string
)

func init() {
	BaseURL = os.Getenv("ELAS_BASE_URL")

	resp, err := http.Get(fmt.Sprintf("%s/config/doc/faq", BaseURL))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	conf := &struct {
		Source map[string]string `json:"_source"`
	}{}
	if err := json.Unmarshal(body, conf); err != nil {
		panic(err)
	}

	UserIndex = conf.Source["user_index"]
	UserType = conf.Source["user_type"]
	SecretKey = conf.Source["auth_secret_key"]
	fmt.Println(conf.Source)
}
