package login

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserInfo は クライアントから渡されたログイン情報を保持する構造体です
type UserInfo struct {
	MailAddress string `json:"mail_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

// VerifyLoginResult は クライアントから渡されたユーザー情報をもとにDBへデータが存在するかの問い合わせをした結果を格納します
type VerifyLoginResult struct {
	UserID    int  `json:"user_id"`
	LoginFlag bool `json:"login_flag"`
}

// Get は users テーブルから user 情報を取得します
func (r UserInfo) Get(ctx context.Context) (models.User, error) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo_app")
	if err != nil {
		log.Print(err)
		return models.User{}, errors.New(consts.SystemError)
	}
	defer db.Close()

	user, err := models.Users(
		qm.Where("user_name=?", r.UserName),
		qm.Where("mail_address=?", r.MailAddress),
	).One(ctx, db)

	if err != nil {
		return models.User{}, err
	}

	return *user, err
}
