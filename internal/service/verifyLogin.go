package service

import (
	"context"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure"
)

// VerifyLogin はログイン可否を判定します
func VerifyLogin(ctx context.Context, loginInfo infrastructure.LoginInfo) (infrastructure.TodoList, error) {

	todoList, err := loginInfo.Get(ctx)
	if err != nil {
		return infrastructure.TodoList{}, err
	}

	// テスト用
	todoList = infrastructure.TodoList{
		UserId:      1,
		TodoId:      1,
		ActiveTask:  "トイレしなきゃ",
		Description: "もれそうです",
	}
	return todoList, nil
}
