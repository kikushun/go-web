package controller

import (
	"log"
	"net/http"
)

// Handler handler
type Handler struct {
	mainProcess func(http.ResponseWriter, *http.Request)
}

// 共通処理はここ
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	h.mainProcess(w, r)
	log.Println("end")
}
