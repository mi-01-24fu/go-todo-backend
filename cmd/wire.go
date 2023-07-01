//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	handlerAddition "github.com/mi-01-24fu/go-todo-backend/internal/handlers/addition"
	handlerGetList "github.com/mi-01-24fu/go-todo-backend/internal/handlers/get_list"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/addition"
	getList "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
	serviceAddition "github.com/mi-01-24fu/go-todo-backend/internal/service/addition"
	serviceGetList "github.com/mi-01-24fu/go-todo-backend/internal/service/get_list"
)

func initializeGetListEvent(db *sql.DB) *handlerGetList.GetListHandler {
	wire.Build(
		getList.NewAccessTODOImpl,
		serviceGetList.NewGetService,
		handlerGetList.NewGetListHandler,
	)
	return &handlerGetList.GetListHandler{}
}

func initializeAdditionEvent(db *sql.DB) *handlerAddition.AdditionImple {
	wire.Build(
		addition.NewAdditionTaskImpl,
		serviceAddition.NewVerifyAdditionImpl,
		handlerAddition.NewAdditionImple,
	)
	return &handlerAddition.AdditionImple{}
}
