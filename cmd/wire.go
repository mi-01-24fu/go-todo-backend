//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	todoHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/todo"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/todo"
	"github.com/mi-01-24fu/go-todo-backend/internal/service/todo"
)

func InitializeEvent(db *sql.DB) *todoHandler.TODOListHandler {
	wire.Build(todoHandler.NewTODOListHandler, todo.NewGetService, access.NewAccessTODOImpl)
	return &todoHandler.TODOListHandler{}
}
