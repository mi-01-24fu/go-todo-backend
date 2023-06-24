package service_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
// 	"github.com/mi-01-24fu/go-todo-backend/internal/service"
// )

// func TestVerifyLogin(t *testing.T) {

// 	// ctrl := gomock.NewController(t)
// 	// defer ctrl.Finish()

// 	// repoMock := login.NewMockAccessLoginInfo(ctrl)
// 	// repoMock.EXPECT().Get(gomock.Any()).Return(models.User{ID: 1, UserName: "mifu", MailAddress: "inogan38@gmail.com"})

// 	type args struct {
// 		loginInfo login.UserInfo
// 	}
// 	tests := []struct {
// 		name    string
// 		ctx     context.Context
// 		args    args
// 		want    login.VerifyLoginResult
// 		wantErr bool
// 	}{
// 		{
// 			"正常/TodoListとnilを返却する",
// 			context.Background(),
// 			args{
// 				login.UserInfo{
// 					MailAddress: "inogan38@gmail.com",
// 					UserName:    "mifu",
// 				},
// 			},
// 			login.VerifyLoginResult{
// 				UserID:    1,
// 				LoginFlag: true,
// 			},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := service.VerifyLogin(tt.ctx, tt.args.loginInfo)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("VerifyLogin() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("VerifyLogin() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
