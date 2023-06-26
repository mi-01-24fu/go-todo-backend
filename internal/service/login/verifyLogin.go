package service

import (
	"context"
	"errors"
	"log"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	login "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
)

// LoginRepository は ログイン可否を判定を行うインターフェース定義です
type LoginRepository interface {
	VerifyLogin(context.Context, login.UserInfo) (login.VerifyLoginResult, error)
}

// VerifyLoginInfo は LoginRepository インターフェースを満たす構造体
type VerifyLoginInfo struct{}

const noRecord = "sql: no rows in result set"

// VerifyLogin はログイン可否を判定します
func (VerifyLoginInfo) VerifyLogin(ctx context.Context, loginInfo login.UserInfo) (login.VerifyLoginResult, error) {

	user, err := loginInfo.Get(ctx)

	if err != nil {
		if err.Error() == noRecord {
			log.Print(err)
			return login.VerifyLoginResult{
				UserID:    0,
				LoginFlag: false,
			}, nil
		}
	}
	if err != nil {
		log.Print(err)
		return login.VerifyLoginResult{}, errors.New(consts.SystemError)
	}

	returnTodoList := login.VerifyLoginResult{
		UserID:    user.ID,
		LoginFlag: true,
	}
	return returnTodoList, nil
}
