package login

import (
	"context"
	"errors"
	"log"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	login "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
)

// LoginRepository は ログイン可否を判定を行うインターフェース定義です
type LoginRepository interface {
	VerifyLogin(context.Context, login.RequestUser) (login.ResponseUser, error)
}

// LoginRepositoryImpl は LoginRepository インターフェースを満たす構造体
type LoginRepositoryImpl struct {
	AccessLogin login.ConfirmLogin
}

// NewLoginRepositoryImpl は新しい LoginRepositoryImpl 構造体を作成して返却する
func NewLoginRepositoryImpl(a login.ConfirmLogin) LoginRepository {
	return &LoginRepositoryImpl{AccessLogin: a}
}

const noRecord = "sql: no rows in result set"

// VerifyLogin はログイン可否を判定します
func (l LoginRepositoryImpl) VerifyLogin(ctx context.Context, requestUser login.RequestUser) (login.ResponseUser, error) {

	user, err := l.AccessLogin.Get(ctx, requestUser)

	if err != nil {
		if err.Error() == noRecord {
			log.Print(err)
			return login.ResponseUser{
				UserID:    0,
				LoginFlag: false,
			}, err
		}
	}
	if err != nil {
		log.Print(err)
		return login.ResponseUser{}, errors.New(consts.SystemError)
	}

	return login.ResponseUser{
		UserID:    user.UserID,
		LoginFlag: true,
	}, nil
}
