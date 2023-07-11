package verifySignup

import (
	"context"
	"math/rand"
	"errors"
	"log"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/configuration"
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
	err = s.AccessSignUpRepo.Count(ctx, requestData.MailAddress)
	if err != nil {
		log.Print(err)
		return verifySignup.VerifySignUpResponse{}, err
	}

	useSESFlag, err := configuration.UseSES()
	if err != nil {
		log.Print(err)
		return verifySignup.VerifySignUpResponse{}, err
	}

	// flag判定 true
	if useSESFlag {
		// 乱数作成

		// 仮会員登録(ユーザー名,mailAddressを登録するがフラグがfalse)
		// ConfirmMail呼出し
		// 乱数とctxとユーザー名、mailAddressを渡す
		// htmlテンプレート作成
		// レスポンス返却(乱数)
		log.Print(err)
	} else {
		// flag判定 false
		// awsを利用しない場合は新規会員登録完了
		// AWS SESを利用しない場合は新規下院登録完了なので、それをフロントに知らせるflagも必要
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
