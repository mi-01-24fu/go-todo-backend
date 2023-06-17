package handlers

import (
	"encoding/json"
	"errors"
	"go-todo-backend/internal/service"
	"net/http"
	"regexp"
)

const regex = `^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`

/*
MailAddress, UserNameを受け取る
dataがあるかの確認を行う
*/
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var loginInfo service.LoginInfo

	if err := json.Unmarshal([]byte(body), &loginInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := checkRequestData(loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := service.VerifyLogin(loginInfo)
}

func checkRequestData(data LoginInfo) error {

	if data.MailAddress == "" {
		return errors.New("MailAddress が入力されていません")
	}

	if !regexp.MustCompile(regex).MatchString(data.MailAddress) {
		return errors.New("MailAddress の形式が正しくありません")
	}

	if data.UserName == "" {
		return errors.New("UserName が入力されていません")
	}

	if len(data.UserName) < 3 && len(data.UserName) > 12 {
		return errors.New("UserName の文字数が不正です")
	}
	return nil
}
