package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mi-01-24fu/go-todo-backend/internal/handlers"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	login "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
	signup "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

type dbConfig struct {
	database     string
	user         string
	password     string
	protocol     string
	host         string
	port         string
	databaseName string
}

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	handleSetUp := setUpDBConfig()

	http.HandleFunc("/login", handleSetUp.loginWire)
	http.HandleFunc("/signUp", handleSetUp.signUp)
	http.ListenAndServe(":8080", nil)
}

// setUpDBConfig は DB 接続するための情報を環境変数ファイルから取得します
func setUpDBConfig() dbConfig {
	return dbConfig{
		database:     os.Getenv("DATABASE"),
		user:         os.Getenv("USER"),
		password:     os.Getenv("PASSWORD"),
		protocol:     os.Getenv("PROTOCOL"),
		host:         os.Getenv("HOST"),
		port:         os.Getenv("PORT"),
		databaseName: os.Getenv("DATABASENAME"),
	}
}

// retrunDBOpenString は DB に接続するための文字列情報を返却します
func (d dbConfig) retrunDBConnectionString() (string, string) {
	databaseName := d.database
	connectionInfo := d.user + ":" + d.password + "@" + d.protocol + "(" + d.host + ":" + d.port + ")/" + d.databaseName
	return databaseName, connectionInfo
}

func (d dbConfig) loginWire(w http.ResponseWriter, req *http.Request) {
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

func (d dbConfig) signUp(w http.ResponseWriter, req *http.Request) {

	databaseName, connectionInfo := d.retrunDBConnectionString()
	db, err := sql.Open(databaseName, connectionInfo)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	accessRepo := access.ServiceImpl{DB: db}
	signUpRepo := signup.AccessInfo{AccessRepo: accessRepo}
	signUpService := handlers.NewSignUpService(signUpRepo, accessRepo)
	result, err := signUpService.SignUp(w, req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
