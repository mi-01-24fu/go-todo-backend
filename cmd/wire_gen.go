// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/mi-01-24fu/go-todo-backend/internal/handlers/addition"
	"github.com/mi-01-24fu/go-todo-backend/internal/handlers/get_list"
	addition2 "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/addition"
	get_list2 "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
	addition3 "github.com/mi-01-24fu/go-todo-backend/internal/service/addition"
	get_list3 "github.com/mi-01-24fu/go-todo-backend/internal/service/get_list"
)

// Injectors from wire.go:

func InitializeGetListEvent(db *sql.DB) *get_list.GetListHandler {
	accessTODOImpl := get_list2.NewAccessTODOImpl(db)
	getService := get_list3.NewGetService(accessTODOImpl)
	getListHandler := get_list.NewGetListHandler(getService)
	return getListHandler
}

func initializeAdditionEvent(db *sql.DB) *addition.AdditionImple {
	additionTaskImpl := addition2.NewAdditionTaskImpl(db)
	verifyAdditionImpl := addition3.NewVerifyAdditionImpl(additionTaskImpl)
	additionImple := addition.NewAdditionImple(verifyAdditionImpl)
	return additionImple
}