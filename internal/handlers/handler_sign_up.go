package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	service "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

// SignUpService は SignUp, AccessSignUp インターフェースを保持する構造体
type SignUpService struct {
	SignUpRepo service.SignUp
	AccessRepo signup.AccessSignUp
}

// NewSignUpService は SignUpService 構造体を返却するコンストラクタ関数
func NewSignUpService(r service.SignUp, a signup.AccessSignUp) *SignUpService {
	return &SignUpService{
		SignUpRepo: r,
		AccessRepo: a,
	}
}

// SignUp はユーザー情報の登録処理をハンドリングするメソッド
func (s *SignUpService) SignUp(w http.ResponseWriter, req *http.Request) (signup.VerifySignUpResult, error) {

	if req.Method != "POST" {
		return signup.VerifySignUpResult{}, errors.New("システム障害：大変申し訳ありませんが、一定時間間隔を空けてログインしてください。")
	}

	// リクエストデータの存在有無チェック
	signupInfo, err := checkSignUpInput(req)
	if err != nil {
		return signup.VerifySignUpResult{}, err
	}

	result, err := s.SignUpRepo.VerifySignUp(signupInfo)
	if err != nil {
		return signup.VerifySignUpResult{}, err
	}
	return signup.VerifySignUpResult{UserID: result.UserID, LoginFlag: result.LoginFlag}, nil
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkSignUpInput(req *http.Request) (signup.NewMemberInfo, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return signup.NewMemberInfo{}, errors.New("入力値に問題があります。入力した内容に誤りが無いか確認してください。")
	}

	var signUpInfo signup.NewMemberInfo

	if err := json.Unmarshal(body, &signUpInfo); err != nil {
		return signup.NewMemberInfo{}, errors.New("入力値に問題があります。入力した内容に誤りが無いか確認してください。")
	}
	return signUpInfo, nil
}