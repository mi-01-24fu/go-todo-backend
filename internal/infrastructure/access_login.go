package infrastructure

import (
	"context"
	"github.com/mi-01-24fu/go-todo-backend/internal/service"
)

type TodoList struct {
	UserId      int    `json:"user_id,omitempty"`
	TodoId      int    `json:"todo_id,omitempty"`
	ActiveTask  string `json:"active_task,omitempty"`
	Description string `json:"description,omitempty"`
}

type AccessLoginInfo interface {
	Get(context.Context) (TodoList, error)
}

func (r LoginInfo) Get(ctx context.Context) (TodoList, error) {}