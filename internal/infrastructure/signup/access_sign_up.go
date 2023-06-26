package signup

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// VerifySignUpResult は クライアントから渡されたユーザー情報をもとにDBへデータを登録した結果を格納します
type VerifySignUpResult struct {
	UserID    int  `json:"user_id"`
	LoginFlag bool `json:"login_flag"`
}

// RegistrationRequest は クライアントから渡されたログイン情報を保持する構造体です
type RegistrationRequest struct {
	UserName    string `json:"user_name,omitempty"`
	MailAddress string `json:"mail_address,omitempty"`
}

// SignUpService は ユーザー情報を登録するためのインターフェース
type SignUpService interface {
	Count(string) (int64, error)
	SignUp(RegistrationRequest) (VerifySignUpResult, error)
}

// SignUpServiceImpl は AccessSignUp を満たす構造体
type SignUpServiceImpl struct {
	DB *sql.DB
}

func (a SignUpServiceImpl) Count(mailaddress string) (int64, error) {
	ctx := context.Background()

	m, err := models.Users(
		models.UserWhere.MailAddress.EQ(mailaddress),
	).Count(ctx, a.DB)

	if err != nil {
		return 0, errors.New(consts.SystemError)
	}

	if m != 0 {
		log.Print(err)
		return 0, errors.New(consts.DuplicationMailAddress)
	}

	return m, nil
}

// SignUp はユーザー情報をDBへ登録する処理を行う
func (a SignUpServiceImpl) SignUp(signUpInfo RegistrationRequest) (VerifySignUpResult, error) {

	ctx := context.Background()

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo_app")
	if err != nil {
		log.Print(err)
		return VerifySignUpResult{}, errors.New(consts.SystemError)
	}
	defer db.Close()

	users := models.User{
		UserName:    signUpInfo.UserName,
		MailAddress: signUpInfo.MailAddress,
	}

	err = users.Insert(ctx, db, boil.Infer())

	if err != nil {
		return VerifySignUpResult{}, err
	}

	return VerifySignUpResult{}, err
}
