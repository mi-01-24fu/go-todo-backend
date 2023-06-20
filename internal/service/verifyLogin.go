package service

import (
	"fmt"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure"
)

type LoginInfo struct {
	MailAddress string `json:"mail_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

// VerifyLogin はログイン可否を判定します
func VerifyLogin(loginInfo LoginInfo) (infrastructure.TodoList, error) {
	todoList := infrastructure.TodoList{
		UserId:      1,
		TodoId:      1,
		ActiveTask:  "トイレしなきゃ",
		Description: "もれそうです",
	}
	fmt.Println("a")

	return todoList, nil
}
