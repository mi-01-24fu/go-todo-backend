package verifySignup

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
	verifySignUp "github.com/mi-01-24fu/go-todo-backend/internal/service/verifySignup"
)

// HandlerSignUpRepo は SignUp, AccessSignUp インターフェースを保持する構造体
type HandlerSignUpRepo struct {
	SingUpServiceRepo verifySignUp.PreparationSingUp
}

// NewHandlerSignUpRepo は新しい HandlerSignUpRepo 構造体を作成して返却する
func NewHandlerSignUpRepo(p verifySignUp.PreparationSingUp) *HandlerSignUpRepo {
	return &HandlerSignUpRepo{
		SingUpServiceRepo: p,
	}
}

// VerifySignUp はユーザー情報の登録処理をハンドリングするメソッド
func (h *HandlerSignUpRepo) VerifySignUp(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	// リクエストデータの存在有無チェック
	signupInfo, err := checkSignUpInput(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	responseData, err := h.SingUpServiceRepo.VerifySignUp(ctx, signupInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	createResponse(w, responseData)
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkSignUpInput(req *http.Request) (verifySignup.VerifySignUpRequest, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return verifySignup.VerifySignUpRequest{}, errors.New(consts.BadInput)
	}

	var requestData verifySignup.VerifySignUpRequest

	if err := json.Unmarshal(body, &requestData); err != nil {
		return verifySignup.VerifySignUpRequest{}, errors.New(consts.BadInput)
	}
	return requestData, nil
}

// createResponse はレスポンスデータを加工して返却する
func createResponse(w http.ResponseWriter, responseData verifySignup.VerifySignUpResponse) {
	res, err := json.Marshal(responseData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
