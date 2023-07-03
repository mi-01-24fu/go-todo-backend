package signup

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport
	"github.com/mi-01-24fu/go-todo-backend/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AccessSignUp は ユーザー情報を登録するためのインターフェース
type AccessSignUp interface {
	SignUp(context.Context, RegistrationRequest) (RegistrationResponse, error)
}

// AccessSignUpImpl は AccessSignUp を満たす構造体
type AccessSignUpImpl struct {
	DB *sql.DB
}

// NewAccessSignUpImpl は新しい AccessSignUpImpl を作成し返却する
func NewAccessSignUpImpl(db *sql.DB) *AccessSignUpImpl {
	return &AccessSignUpImpl{DB: db}
}

// RegistrationRequest は クライアントから渡されたログイン情報を保持する構造体です
type RegistrationRequest struct {
	UserName    string `json:"user_name,omitempty"`
	MailAddress string `json:"mail_address,omitempty"`
}

// RegistrationResponse は クライアントから渡されたユーザー情報をもとにDBへデータを登録した結果を格納します
type RegistrationResponse struct {
	UserID    int  `json:"user_id"`
	LoginFlag bool `json:"login_flag"`
}

// SignUp はユーザー情報をDBへ登録する処理を行う
func (a AccessSignUpImpl) SignUp(ctx context.Context, RequestData RegistrationRequest) (RegistrationResponse, error) {

	users := models.User{
		UserName:    RequestData.UserName,
		MailAddress: RequestData.MailAddress,
	}

	err := users.Insert(ctx, a.DB, boil.Infer())

	if err != nil {
		return RegistrationResponse{}, err
	}

	return RegistrationResponse{}, err
}
