package main

import (
	"encoding/json"
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/handlers"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	login "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
	signup "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

func main() {
	http.HandleFunc("/login", loginWire)
	http.HandleFunc("/signUp", signUp)
	http.ListenAndServe(":8080", nil)
}

func loginWire(w http.ResponseWriter, req *http.Request) {
	verifyLoginInfo := login.VerifyLoginInfo{}
	loginService := handlers.Service{Repo: verifyLoginInfo}

	result, err := loginService.LoginHandler(w, req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	accessRepo := access.AccessSignUpInfo{}
	signUpRepo := signup.AccessInfo{AccessRepo: accessRepo}
	signUpService := handlers.NewSignUpService(signUpRepo, accessRepo)
	signUpService.SignUp(w, req)
}
