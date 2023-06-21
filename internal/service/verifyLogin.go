package service

import (
	"context"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
)

// VerifyLogin はログイン可否を判定します
func VerifyLogin(ctx context.Context, loginInfo login.LoginInfo) (login.TodoList, error) {

	todoList, err := loginInfo.Get(ctx)
	if err != nil {
		return login.TodoList{}, err
	}

	// テスト用
	todoList = login.TodoList{
		UserId:      1,
		TodoId:      1,
		ActiveTask:  "トイレしなきゃ",
		Description: "もれそうです",
	}
	return todoList, nil
}
