package service

import (
	"context"
	"errors"
	"fmt"

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
			fmt.Println(err)
			return login.VerifyLoginResult{
				UserID:    0,
				LoginFlag: false,
			}, nil
		}
	}
	if err != nil {
		fmt.Println(err)
		return login.VerifyLoginResult{}, errors.New("システム障害：大変申し訳ありませんが、一定時間間隔を空けてログインしてください。")
	}

	returnTodoList := login.VerifyLoginResult{
		UserID:    user.ID,
		LoginFlag: true,
	}
	return returnTodoList, nil
}
