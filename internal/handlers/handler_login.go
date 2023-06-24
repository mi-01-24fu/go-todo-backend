package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	service "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
)

// Service は LoginRepository インターフェースを保持する構造体
type Service struct {
	Repo service.LoginRepository
}

// LoginHandler は login リクエストを受けてログインが可能かの判定結果を返します
func (s Service) LoginHandler(w http.ResponseWriter, req *http.Request) (login.VerifyLoginResult, error) {

	if req.Method != "POST" {
		return login.VerifyLoginResult{}, errors.New("システム障害：大変申し訳ありませんが、一定時間間隔を空けてログインしてください。")
	}

	// リクエストデータの存在有無チェック
	loginInfo, err := checkRequestData(req)
	if err != nil {
		return login.VerifyLoginResult{}, err
	}

	// リクエストデータのバリデーションチェック
	err = checkValidation(loginInfo)
	if err != nil {
		return login.VerifyLoginResult{}, err
	}

	ctx := context.Background()

	// ログイン判定処理呼出し
	result, err := s.Repo.VerifyLogin(ctx, loginInfo)

	return result, err
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkRequestData(req *http.Request) (login.UserInfo, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return login.UserInfo{}, errors.New("入力値に問題があります。入力した内容に誤りが無いか確認してください。")
	}

	var loginInfo login.UserInfo

	if err := json.Unmarshal(body, &loginInfo); err != nil {
		return login.UserInfo{}, errors.New("入力値に問題があります。入力した内容に誤りが無いか確認してください。")
	}
	return loginInfo, nil
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data login.UserInfo) error {
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
