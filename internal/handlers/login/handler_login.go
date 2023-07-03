package login

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Println("--1--", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("--2--")

	// リクエストデータのバリデーションチェック
	err = checkValidation(loginInfo)
	fmt.Println("--3--")
	if err != nil {
		fmt.Println("--4--")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	fmt.Println("--5--")
	// ログイン判定処理呼出し
	responseData, err := s.LoginRepo.VerifyLogin(ctx, loginInfo)
	if err != nil {
		fmt.Println("--6--")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンス返却
	createResponse(w, responseData)
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkRequestData(req *http.Request) (login.RequestUser, error) {

	fmt.Println("--a--")
	body, err := ioutil.ReadAll(req.Body)
	fmt.Println("--b--")
	if err != nil {
		fmt.Println("--c--")
		return login.RequestUser{}, errors.New(consts.BadInput)
	}

	var loginInfo login.RequestUser

	fmt.Println("aaa", string(body))

	if err := json.Unmarshal(body, &loginInfo); err != nil {
		fmt.Println("--d--")
		fmt.Println("--d--", err)
		return login.RequestUser{}, errors.New(consts.BadInput)
	}
	fmt.Println("--e--")
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
