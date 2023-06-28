package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	loginHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/login"
	signupHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/signup"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	login "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
	signup "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

type dbConfig struct {
	databaseType string
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

	// DB へ接続する情報を環境変数ファイルから取得
	dbSettings := setUpDBConfig()

	// DB へ接続するための文字列を加工
	databaseName, connectionInfo := dbSettings.generateDBConnectionString()

	// DB接続
	db, err := databaseOpen(databaseName, connectionInfo)
	if err != nil {
		panic("システムエラー")
	}
	defer db.Close()

	// ---結局のところInitializeEventでやっていることは同じ---
	// databaseSetup := getTODO.AccessTODOImpl{DB: db}
	// getService := todo.GetService{AccessRepository: databaseSetup}
	// getHandler := handlers.GetTODOService{GetTODORepo: getService}

	// 構造体の初期化
	//event := InitializeEvent(db)

	http.HandleFunc("/login", dbSettings.loginWire)
	http.HandleFunc("/signUp", dbSettings.signUp)
	//http.HandleFunc("/getTODOList", event.GetTODOList)
	http.ListenAndServe(":8080", nil)
}

// setUpDBConfig は DB 接続するための情報を環境変数ファイルから取得します
func setUpDBConfig() dbConfig {
	return dbConfig{
		databaseType: os.Getenv("DATABASE"),
		user:         os.Getenv("USER"),
		password:     os.Getenv("PASSWORD"),
		protocol:     os.Getenv("PROTOCOL"),
		host:         os.Getenv("HOST"),
		port:         os.Getenv("PORT"),
		databaseName: os.Getenv("DATABASENAME"),
	}
}

// retrunDBOpenString は DB に接続するための文字列情報を返却します
func (d dbConfig) generateDBConnectionString() (string, string) {
	databaseName := d.databaseType
	connectionInfo := d.user + ":" + d.password + "@" + d.protocol + "(" + d.host + ":" + d.port + ")/" + d.databaseName
	return databaseName, connectionInfo
}

func (d dbConfig) loginWire(w http.ResponseWriter, req *http.Request) {
	verifyLoginInfo := login.VerifyLoginInfo{}
	loginService := loginHandler.Service{Repo: verifyLoginInfo}

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

// DBへ接続する
func databaseOpen(databaseName, connectionInfo string) (*sql.DB, error) {
	db, err := sql.Open(databaseName, connectionInfo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return db, nil
}

func (d dbConfig) signUp(w http.ResponseWriter, req *http.Request) {

	databaseName, connectionInfo := d.generateDBConnectionString()
	db, err := sql.Open(databaseName, connectionInfo)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	accessRepo := access.ServiceImpl{DB: db}
	signUpRepo := signup.AccessInfo{AccessRepo: accessRepo}
	signUpService := signupHandler.NewSignUpService(signUpRepo, accessRepo)
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
