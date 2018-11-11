package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Login login
func Login(w *http.ResponseWriter, r *http.Request) interface{} {
	id := r.PostFormValue("id")
	password := r.PostFormValue("password")
	fmt.Println(password)
	// elasのユーザに存在するかチェック

	// success
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(1 * time.Hour),
	})

	tokenStr, err := token.SignedString("secretKey")
	if err != nil {
		return err
	}
	fmt.Println(tokenStr)
	return nil
}
