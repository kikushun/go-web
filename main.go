package main

import (
	"net/http"

	"github.com/kikuchi/go-web/controller"
)

func main() {
	mux := http.NewServeMux()

	// ログイン
	mux.HandleFunc("/login", controller.Login)
	// カテゴリ
	mux.Handle("/category", h("GET", controller.GetCategory))
	mux.Handle("/category/search", h("GET", controller.SearchCategory))
	// ユーザー
	mux.Handle("/user/save", h("POST", controller.SaveUser))
	mux.Handle("/user/search", h("GET", controller.SearchUser))
	mux.Handle("/user/delete", h("GET", controller.DeleteUser))

	http.ListenAndServe(":9090", mux)
}
