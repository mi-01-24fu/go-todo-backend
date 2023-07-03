package login

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	login "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
)

func TestVerifyLogin(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestUser := login.RequestUser{
		MailAddress: "inogan38@gmail.com",
		UserName:    "mifu",
	}

	const noRecord = "sql: no rows in result set"

	responseUser := access.ResponseUser{
		UserID:    1,
		LoginFlag: true,
	}

	type args struct {
		loginInfo access.RequestUser
	}
	tests := []struct {
		name        string
		setup       func(*access.MockConfirmLogin)
		ctx         context.Context
		args        args
		want        access.ResponseUser
		errorString string
		wantErr     bool
	}{
		{
			"正常/TodoListとnilを返却する",
			func(mcl *access.MockConfirmLogin) {
				mcl.EXPECT().Get(gomock.Any(), requestUser).Return(responseUser, nil)
			},
			context.Background(),
			args{
				requestUser,
			},
			responseUser,
			"",
			false,
		},
		{
			"準正常/noRecordエラー",
			func(mcl *access.MockConfirmLogin) {
				mcl.EXPECT().Get(gomock.Any(), requestUser).Return(login.ResponseUser{}, errors.New(noRecord))
			},
			context.Background(),
			args{
				requestUser,
			},
			login.ResponseUser{},
			noRecord,
			true,
		},
		{
			"異常/Get呼出しエラー",
			func(mcl *access.MockConfirmLogin) {
				mcl.EXPECT().Get(gomock.Any(), requestUser).Return(login.ResponseUser{}, errors.New(consts.SystemError))
			},
			context.Background(),
			args{
				requestUser,
			},
			login.ResponseUser{},
			consts.SystemError,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockConfirmLogin := access.NewMockConfirmLogin(ctrl)
			if tt.setup != nil {
				tt.setup(mockConfirmLogin)
			}
			l := LoginRepositoryImpl{mockConfirmLogin}

			got, err := l.VerifyLogin(tt.ctx, tt.args.loginInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
