package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kikuchi/go-web/config"
	"github.com/kikuchi/go-web/model"
)

// CreateToken 認証成功したユーザのトークン作成
func CreateToken(user *model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_name":  user.Name,
		"user_email": user.Email,
		"exp":        time.Now().Add(1 * time.Minute).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, err
}

// Authenticate 認証
func Authenticate(r *http.Request) error {

	strArr := strings.Fields(r.Header.Get("Authorization"))

	if len(strArr) != 2 {
		return errors.New("[認証エラー] Authorizationヘッダーの文字列が不正です。")
	}

	token, err := jwt.Parse(strArr[1], checkToken)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("[認証エラー] パリデーションエラー")
	}

	return nil
}

// checkToken JWTトークンチェック
func checkToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("[認証エラー] JWTトークンチェック中のエラー")
	}
	return []byte(config.SecretKey), nil
}
