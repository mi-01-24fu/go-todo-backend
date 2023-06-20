package service_test
// package service

// import (
// 	"testing"

// 	getTodoList "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure"
// )

// func TestVerifyLogin(t *testing.T) {
// 	type args struct {
// 		loginInfo LoginInfo
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    getTodoList.TodoList
// 		wantErr bool
// 	}{
// 		{
// 			"正常/TodoListとnilを返却する",
// 			args{
// 				LoginInfo{
// 					MailAddress: "inogan@gmail.com",
// 					UserName:    "mifu",
// 				},
// 			},
// 			getTodoList.TodoList{
// 				UserId:      1,
// 				TodoId:      1,
// 				ActiveTask:  "トイレしなきゃ",
// 				Description: "もれそうです",
// 			},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := VerifyLogin(tt.args.loginInfo)
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
