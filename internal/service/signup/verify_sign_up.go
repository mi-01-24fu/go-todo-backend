package signup

import (
	"errors"
	"log"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
)

// SignUp は VerifySignUp のインターフェース
type SignUp interface {
	VerifySignUp(signup.RegistrationRequest) (signup.VerifySignUpResult, error)
}

// AccessInfo は SignUpインターフェースを満たす構造体
type AccessInfo struct {
	AccessRepo signup.Service
}

// VerifySignUp は 渡されたユーザー情報をもとにDB登録するための準備を行う
func (s AccessInfo) VerifySignUp(signupInfo signup.RegistrationRequest) (signup.VerifySignUpResult, error) {

	// バリデーションチェック
	err := checkValidation(signupInfo)
	if err != nil {
		return signup.VerifySignUpResult{}, err
	}

	// メールアドレスの重複確認
	count, err := s.AccessRepo.Count(signupInfo.MailAddress)
	log.Print(count)
	if err != nil {
		log.Print(err)
		return signup.VerifySignUpResult{}, err
	}

	// 新規会員登録処理呼出し
	result, err := s.AccessRepo.SignUp(signupInfo)
	if err != nil {
		log.Print(err)
		return signup.VerifySignUpResult{}, errors.New(consts.SystemError)
	}
	return result, err
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data signup.RegistrationRequest) error {

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
