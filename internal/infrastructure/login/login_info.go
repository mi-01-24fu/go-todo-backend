package login

import (
	"context"

	"github.com/mi-01-24fu/go-todo-backend/models"
)

// AccessLoginInfo は users テーブル へaccessするインターフェースです
type AccessLoginInfo interface {
	Get(context.Context) (models.User, error)
}
