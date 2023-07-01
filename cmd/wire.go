//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	handlerAddition "github.com/mi-01-24fu/go-todo-backend/internal/handlers/addition"
	handlerGetList "github.com/mi-01-24fu/go-todo-backend/internal/handlers/getlist"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/addition"
	getList "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/getlist"
	serviceAddition "github.com/mi-01-24fu/go-todo-backend/internal/service/addition"
	serviceGetList "github.com/mi-01-24fu/go-todo-backend/internal/service/getlist"
)

func initializeGetListEvent(db *sql.DB) *handlerGetList.TODOGetHandler {
	wire.Build(
		getList.NewAccessTODOImpl,
		serviceGetList.NewGetService,
		handlerGetList.NewGetListHandler,
	)
	return &handlerGetList.TODOGetHandler{}
}

func initializeAdditionEvent(db *sql.DB) *handlerAddition.TaskAdditionImpl {
	wire.Build(
		addition.NewAdditionTaskImpl,
		serviceAddition.NewVerifyAdditionImpl,
		handlerAddition.NewAdditionImple,
	)
	return &handlerAddition.TaskAdditionImpl{}
}
