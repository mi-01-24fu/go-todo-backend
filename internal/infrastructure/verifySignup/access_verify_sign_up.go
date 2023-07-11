package verifySignup

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql" // init関数を実行するためにimport
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/models"
)

// AccessVerifySignUp は ユーザー情報を登録するための検証を行うインターフェース
type AccessVerifySignUp interface {
	Count(context.Context, string) error
	TemporaryStore(context.Context, VerifySignUpRequest) (error, VerifySignUpResponse)
	ConfirmMail(context.Context, string) error
}

// AccessVerifySignUpImpl は AccessSignUp を満たす構造体
type AccessVerifySignUpImpl struct {
	DB *sql.DB
}

// NewVerifySignUpServiceImpl は新しい VerifySignUpServiceImpl を作成し返却する
func NewVerifySignUpServiceImpl(db *sql.DB) *AccessVerifySignUpImpl {
	return &AccessVerifySignUpImpl{DB: db}
}

// VerifySignUpRequest は クライアントから渡されたログイン情報を保持する構造体です
type VerifySignUpRequest struct {
	UserName    string `json:"user_name,omitempty"`
	MailAddress string `json:"mail_address,omitempty"`
}

// VerifySignUpResponse は クライアントから渡されたユーザー情報をもとにDBへデータを登録した結果を格納します
type VerifySignUpResponse struct {
	AuthenticationNumber int
}

// Count は新規会員登録者のメールアドレスが既に登録済みかを確認します
func (a AccessVerifySignUpImpl) Count(ctx context.Context, mailaddress string) error {

	m, err := models.Users(
		models.UserWhere.MailAddress.EQ(mailaddress),
	).Count(ctx, a.DB)

	if err != nil {
		return errors.New(consts.SystemError)
	}

	if m != 0 {
		log.Print(err)
		return errors.New(consts.DuplicationMailAddress)
	}

	return nil
}

func (a AccessVerifySignUpImpl) TempraryStore(ctx context.Context, requestData VerifySignUpRequest) (error, VerifySignUpResponse) {
	const (
		min = 1000
		max = 10000
	)
	
	randomNum := rand.Intn(max-min) + min

	return nil, VerifySignUpResponse{}
}

func (a AccessVerifySignUpImpl) ConfirmMail(ctx context.Context, mailAddress string) error {
	return nil
}
