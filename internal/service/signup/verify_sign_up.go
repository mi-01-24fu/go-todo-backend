package signup

import (
	"errors"

	commonMailAddress "github.com/mi-01-24fu/go-todo-backend/internal/common/mailaddress"
	commonUserName "github.com/mi-01-24fu/go-todo-backend/internal/common/username"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
)

// SignUp は VerifySignUp のインターフェース
type SignUp interface {
	VerifySignUp(signup.NewMemberInfo) (signup.VerifySignUpResult, error)
}

// AccessInfo は SignUpインターフェースを満たす構造体
type AccessInfo struct {
	AccessRepo signup.AccessSignUp
}

// VerifySignUp は 渡されたユーザー情報をもとにDB登録するための準備を行う
func (s AccessInfo) VerifySignUp(signupInfo signup.NewMemberInfo) (signup.VerifySignUpResult, error) {

	err := checkValidation(signupInfo)
	if err != nil {
		return signup.VerifySignUpResult{}, err
	}
	result, err := s.AccessRepo.SignUp(signupInfo)
	return result, err
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data signup.NewMemberInfo) error {

	if !commonUserName.IsEmpty(data.UserName) {
		return errors.New("UserName が入力されていません")
	}

	if !commonUserName.CheckLength(data.UserName) {
		return errors.New("UserName の文字数が不正です")
	}

	if !commonMailAddress.IsEmpty(data.MailAddress) {
		return errors.New("MailAddress が入力されていません")
	}

	if !commonMailAddress.CheckValidation(data.MailAddress) {
		return errors.New("MailAddress の形式が正しくありません")
	}
	return nil
}
