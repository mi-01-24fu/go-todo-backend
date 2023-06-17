package main

import (
	"fmt"
	"go-todo-backend/internal/handlers"
	"net/http"
)

func main() {
	fmt.Println("1")
	http.HandleFunc("/login", handlers.LoginHandler)
	fmt.Println("2")
	http.ListenAndServe(":8080", nil)
}
