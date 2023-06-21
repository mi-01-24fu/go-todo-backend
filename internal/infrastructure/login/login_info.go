package login

import (
	"context"
)

type AccessLoginInfo interface {
	Get(context.Context) (TodoList, error)
}
