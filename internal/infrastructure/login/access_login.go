package login

import (
	"context"
)

type LoginInfo struct {
	MailAddress string `json:"mail_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

type TodoList struct {
	UserId      int    `json:"user_id,omitempty"`
	TodoId      int    `json:"todo_id,omitempty"`
	ActiveTask  string `json:"active_task,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r LoginInfo) Get(ctx context.Context) (TodoList, error) {
	return TodoList{}, nil
}
