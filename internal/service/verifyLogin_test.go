package service_test

import (
	"context"
	"testing"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	"github.com/mi-01-24fu/go-todo-backend/internal/service"
)

func TestVerifyLogin(t *testing.T) {
	type args struct {
		loginInfo login.LoginInfo
	}
	tests := []struct {
		name    string
		ctx     context.Context
		args    args
		want    login.TodoList
		wantErr bool
	}{
		{
			"正常/TodoListとnilを返却する",
			context.Background(),
			args{
				login.LoginInfo{
					MailAddress: "inogan@gmail.com",
					UserName:    "mifu",
				},
			},
			login.TodoList{
				UserId:      1,
				TodoId:      1,
				ActiveTask:  "トイレしなきゃ",
				Description: "もれそうです",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.VerifyLogin(tt.ctx, tt.args.loginInfo)
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
