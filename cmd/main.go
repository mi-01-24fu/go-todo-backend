package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport

	"github.com/joho/godotenv"
	handlerAddition "github.com/mi-01-24fu/go-todo-backend/internal/handlers/addition"
	handlerGetList "github.com/mi-01-24fu/go-todo-backend/internal/handlers/getlist"
	loginHandler "github.com/mi-01-24fu/go-todo-backend/internal/handlers/login"
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
	loginEvent *loginHandler.VerifyLoginHandler
	getEvent   *handlerGetList.TODOGetHandler
	addEvent   *handlerAddition.TaskAdditionImpl
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

	http.HandleFunc("/login", event.loginEvent.LoginHandler)
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

// DBへ接続する
func dbOpen(dbName, connectionInfo string) (*sql.DB, error) {
	db, err := sql.Open(dbName, connectionInfo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return db, nil
}

// initializeEvent 各構造体の初期化を行う
func initializeEvent(db *sql.DB) *event {
	loginEvent := initializeLoginEvent(db)
	getEvent := initializeGetListEvent(db)
	addEvent := initializeAdditionEvent(db)
	return &event{
		loginEvent: loginEvent,
		getEvent:   getEvent,
		addEvent:   addEvent,
	}
}
