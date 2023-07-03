package verifySignup

import (
	"context"
	"errors"
	"log"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
)

// Preparation は VerifySignUp のインターフェース
type PreparationSingUp interface {
	VerifySignUp(context.Context, verifySignup.VerifySignUpRequest) (verifySignup.VerifySignUpResponse, error)
}

// PreparationSingUpImpl は SignUpインターフェースを満たす構造体
type PreparationSingUpImpl struct {
	AccessSignUpRepo verifySignup.AccessVerifySignUp
}

// NewPreparationSingUpImpl は新しい PreparationSingUpImpl 構造体を作成して返却する
func NewPreparationSingUpImpl(v verifySignup.AccessVerifySignUp) *PreparationSingUpImpl {
	return &PreparationSingUpImpl{
		AccessSignUpRepo: v,
	}
}

// VerifySignUp は 渡されたユーザー情報をもとにDB登録するための準備を行う
func (s PreparationSingUpImpl) VerifySignUp(ctx context.Context, requestData verifySignup.VerifySignUpRequest) (verifySignup.VerifySignUpResponse, error) {

	// バリデーションチェック
	err := checkValidation(requestData)
	if err != nil {
		return verifySignup.VerifySignUpResponse{}, err
	}

	// メールアドレスの重複確認
	count, err := s.AccessSignUpRepo.Count(ctx, requestData.MailAddress)
	log.Print(count)
	if err != nil {
		log.Print(err)
		return verifySignup.VerifySignUpResponse{}, err
	}
	return verifySignup.VerifySignUpResponse{}, nil
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data verifySignup.VerifySignUpRequest) error {

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
