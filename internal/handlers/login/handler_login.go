package login

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	service "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
)

// VerifyLoginHandler は LoginRepository インターフェースを保持する構造体
type VerifyLoginHandler struct {
	LoginRepo service.LoginRepository
}

// NewVerifyLoginHandler は新しい VerifyLoginHandler 構造体を作成して返却する
func NewVerifyLoginHandler(s service.LoginRepository) *VerifyLoginHandler {
	return &VerifyLoginHandler{LoginRepo: s}
}

// LoginHandler は login リクエストを受けてログインが可能かの判定結果を返します
func (s VerifyLoginHandler) LoginHandler(w http.ResponseWriter, req *http.Request) {

	// リクエストデータの存在有無チェック
	loginInfo, err := checkRequestData(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// リクエストデータのバリデーションチェック
	err = checkValidation(loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := context.Background()

	// ログイン判定処理呼出し
	responseData, err := s.LoginRepo.VerifyLogin(ctx, loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// レスポンス返却
	createResponse(w, responseData)
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkRequestData(req *http.Request) (login.RequestUser, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return login.RequestUser{}, errors.New(consts.BadInput)
	}

	var loginInfo login.RequestUser

	if err := json.Unmarshal(body, &loginInfo); err != nil {
		return login.RequestUser{}, errors.New(consts.BadInput)
	}
	return loginInfo, nil
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data login.RequestUser) error {

	if !commonUserName.IsEmpty(data.UserName) {
		return errors.New(consts.EmptyUserName)
	}

	if !commonUserName.CheckLength(data.UserName) {
		return errors.New(consts.InvalidUserNameLength)
	}

	if !commonMailAddress.IsEmpty(data.MailAddress) {
		return errors.New(consts.EmptyMailAddress)
	}

	if !commonMailAddress.CheckValidation(data.MailAddress) {
		return errors.New(consts.NotMailAddressFormat)
	}

	return nil
}

// createResponse はレスポンスデータを加工して返却する
func createResponse(w http.ResponseWriter, data login.ResponseUser) {
	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
