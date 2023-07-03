package verifySignup

import (
	"context"
	"errors"
	"log"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
)

// ServiceSingUp は VerifySignUp のインターフェース
type ServiceSingUp interface {
	SignUp(context.Context, signup.RegistrationRequest) (signup.RegistrationResponse, error)
}

// ServiceSingUpImpl は SignUpインターフェースを満たす構造体
type ServiceSingUpImpl struct {
	AccessSignUpRepo signup.AccessSignUp
}

// NewServiceSingUpImpl は新しい PreparationSingUpImpl 構造体を作成して返却する
func NewServiceSingUpImpl(a signup.AccessSignUp) *ServiceSingUpImpl {
	return &ServiceSingUpImpl{
		AccessSignUpRepo: a,
	}
}

// VerifySignUp は 渡されたユーザー情報をもとにDB登録するための準備を行う
func (s ServiceSingUpImpl) VerifySignUp(ctx context.Context, requestData signup.RegistrationRequest) (signup.RegistrationResponse, error) {

	// 新規会員登録処理呼出し
	result, err := s.AccessSignUpRepo.SignUp(ctx, requestData)
	if err != nil {
		log.Print(err)
		return signup.RegistrationResponse{}, errors.New(consts.SystemError)
	}
	return result, nil
}
