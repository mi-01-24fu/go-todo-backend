package verifySignup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// SignUp はユーザー情報の登録処理をハンドリングするメソッド
func (h *HandlerSignUpRepo) SignUp(w http.ResponseWriter, req *http.Request) (verifySignup.VerifySignUpResponse, error) {

	ctx := context.Background()

	// リクエストデータの存在有無チェック
	signupInfo, err := checkSignUpInput(req)
	if err != nil {
		return verifySignup.VerifySignUpResponse{}, err
	}

	result, err := h.SingUpServiceRepo.VerifySignUp(ctx, signupInfo)
	fmt.Println(result)
	if err != nil {
		return verifySignup.VerifySignUpResponse{}, err
	}
	return verifySignup.VerifySignUpResponse{}, nil
}

// checkRequestData は望むリクエストデータが送られてきているかを確認します
func checkSignUpInput(req *http.Request) (verifySignup.VerifySignUpRequest, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return verifySignup.VerifySignUpRequest{}, errors.New(consts.BadInput)
	}

	var signUpInfo verifySignup.VerifySignUpRequest

	if err := json.Unmarshal(body, &signUpInfo); err != nil {
		return verifySignup.VerifySignUpRequest{}, errors.New(consts.BadInput)
	}
	return signUpInfo, nil
}
