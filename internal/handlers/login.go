package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mail_address"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/user_name"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure"
	"github.com/mi-01-24fu/go-todo-backend/internal/service"
)

// LoginHandler は login リクエストを受けてログインが可能かの判定結果を返します
func LoginHandler(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	// リクエストデータの存在有無チェック
	loginInfo, err := checkRequestData(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// リクエストデータのバリデーションチェック
	err = checkValidation(loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ログイン判定処理呼出し
	result, err := service.VerifyLogin(ctx, loginInfo)
	fmt.Println(result)
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkRequestData(req *http.Request) (infrastructure.LoginInfo, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return infrastructure.LoginInfo{}, errors.New("リクエストデータを読み込めません: " + err.Error())
	}

	var loginInfo infrastructure.LoginInfo

	if err := json.Unmarshal(body, &loginInfo); err != nil {
		return infrastructure.LoginInfo{}, errors.New("json形式から構造体への変換に失敗しました")
	}
	return loginInfo, nil
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data infrastructure.LoginInfo) error {

	if !commonMailAddress.IsEmpty(data.MailAddress) {
		return errors.New("MailAddress が入力されていません")
	}

	if !commonMailAddress.CheckValidation(data.MailAddress) {
		return errors.New("MailAddress の形式が正しくありません")
	}

	if !commonUserName.IsEmpty(data.UserName) {
		return errors.New("UserName が入力されていません")
	}

	if !commonUserName.CheckLength(data.UserName) {
		return errors.New("UserName の文字数が不正です")
	}
	return nil
}
