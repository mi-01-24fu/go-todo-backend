package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport

	"github.com/joho/godotenv"
	handlerAddition "github.com/mi-01-24fu/go-todo-backend/internal/handlers/addition"
	handlerGetList "github.com/mi-01-24fu/go-todo-backend/internal/handlers/getlist"
	loginHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/login"
	signupHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/signup"

	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	login "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
	signup "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

// dbConfig は database に関する情報を保持する構造体
type dbConfig struct {
	dbType   string
	user     string
	password string
	protocol string
	host     string
	port     string
	dbName   string
}

// event は各機能の依存関係を管理する構造体
type event struct {
	getEvent *handlerGetList.TODOGetHandler
	addEvent *handlerAddition.TaskAdditionImpl
}

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// db接続処理呼出し
	db := dbConnection()
	defer db.Close()

	// 各構造体の初期化
	event := initializeEvent(db)

	// http.HandleFunc("/login", dbSettings.loginWire)
	// http.HandleFunc("/signUp", dbSettings.signUp)
	http.HandleFunc("/getList", event.getEvent.GetTODOList)
	http.HandleFunc("/addition", event.addEvent.TaskAddition)
	http.ListenAndServe(":8080", nil)
}

// dbConnection は database への接続を行う
func dbConnection() *sql.DB {

	// DB へ接続する情報を環境変数ファイルから取得
	dbSettings := setUpDBConfig()

	// DB へ接続するための文字列を加工
	dbName, connectionInfo := dbSettings.generateDBConnectionString()

	// DB接続
	fmt.Println(dbName)
	fmt.Println(connectionInfo)
	db, err := dbOpen(dbName, connectionInfo)
	if err != nil {
		panic("システムエラー")
	}
	return db
}

// setUpDBConfig は DB 接続するための情報を環境変数ファイルから取得します
func setUpDBConfig() dbConfig {
	return dbConfig{
		dbType:   os.Getenv("DBType"),
		user:     os.Getenv("USER"),
		password: os.Getenv("PASSWORD"),
		protocol: os.Getenv("PROTOCOL"),
		host:     os.Getenv("HOST"),
		port:     os.Getenv("PORT"),
		dbName:   os.Getenv("DBNAME"),
	}
}

// retrunDBOpenString は DB に接続するための文字列情報を返却します
func (d dbConfig) generateDBConnectionString() (string, string) {
	dbName := d.dbType
	connectionInfo := d.user + ":" + d.password + "@" + d.protocol + "(" + d.host + ":" + d.port + ")/" + d.dbName
	return dbName, connectionInfo
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
func dbOpen(dbName, connectionInfo string) (*sql.DB, error) {
	fmt.Println(dbName)
	fmt.Println(connectionInfo)
	db, err := sql.Open(dbName, connectionInfo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return db, nil
}

func (d dbConfig) signUp(w http.ResponseWriter, req *http.Request) {

	dbName, connectionInfo := d.generateDBConnectionString()
	db, err := sql.Open(dbName, connectionInfo)
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

// initializeEvent 各構造体の初期化を行う
func initializeEvent(db *sql.DB) *event {
	getEvent := initializeGetListEvent(db)
	addEvent := initializeAdditionEvent(db)
	return &event{
		getEvent: getEvent,
		addEvent: addEvent,
	}
}
