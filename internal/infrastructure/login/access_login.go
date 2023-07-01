package login

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport

	"github.com/mi-01-24fu/go-todo-backend/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ConfirmLogin は ログインを行うためのインターフェース
type ConfirmLogin interface {
	Get(context.Context, RequestUser) (ResponseUser, error)
}

// RequestUser は クライアントから渡されたログイン情報を保持する構造体
type RequestUser struct {
	MailAddress string `json:"mail_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

// ResponseUser は クライアントから渡されたユーザー情報をもとにDBへデータが存在するかの問い合わせをした結果を格納します
type ResponseUser struct {
	UserID    int  `json:"user_id"`
	LoginFlag bool `json:"login_flag"`
}

// ConfirmLoginImpl は ConfirmLogin を満たす構造体
type ConfirmLoginImpl struct {
	DB *sql.DB
}

// NewConfirmLoginImpl は新しい ConfirmLoginImpl 構造体を作成して返却する
func NewConfirmLoginImpl(db *sql.DB) ConfirmLogin {
	return &ConfirmLoginImpl{DB: db}
}

// Get は users テーブルから user 情報を取得します
func (r ConfirmLoginImpl) Get(ctx context.Context, reqestData RequestUser) (ResponseUser, error) {

	user, err := models.Users(
		qm.Where("user_name=?", reqestData.UserName),
		qm.Where("mail_address=?", reqestData.MailAddress),
	).One(ctx, r.DB)
	if err != nil {
		log.Print(err)
		return ResponseUser{}, err
	}

	loginResult := ResponseUser{
		UserID:    user.ID,
		LoginFlag: true,
	}

	return loginResult, err
}
