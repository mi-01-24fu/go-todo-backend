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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AccessVerifySignUp は ユーザー情報を登録するための検証を行うインターフェース
type AccessVerifySignUp interface {
	Count(context.Context, string) error
	TemporaryStore(context.Context, VerifySignUpRequest) (VerifySignUpResponse, error)
	ConfirmMail(context.Context, int, VerifySignUpRequest) error
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

// ConfirmMailInfo は確認用メールを送るための情報を格納する構造体です
type ConfirmMailInfo struct {
	UserName        string
	VerifyNumber    int
	ToMailAddress   string
	FromMailAddress string
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

// TempraryStore はユーザーの仮会員登録を行う
func (a AccessVerifySignUpImpl) TempraryStore(ctx context.Context, requestData VerifySignUpRequest) (VerifySignUpResponse, error) {
	const (
		min = 1000
		max = 10000
	)
	// 検証用整数の設定
	verifyNumber := rand.Intn(max-min) + min

	user := models.User{
		UserName:     requestData.UserName,
		MailAddress:  requestData.MailAddress,
		SignupFlag:   "0", // 0:仮登録完了 1:登録完了
		VerifyNumber: null.IntFrom(verifyNumber),
	}

	err := user.Insert(ctx, a.DB, boil.Infer())
	if err != nil {
		log.Print(err)
		return VerifySignUpResponse{}, err
	}

	return VerifySignUpResponse{
		AuthenticationNumber: verifyNumber,
	}, nil
}

// ConfirmMail はユーザーの MailAddress が正しいかを検証するために AWS SES を利用して確認用メールを送る
func (a AccessVerifySignUpImpl) ConfirmMail(ctx context.Context, mailAddress string) error {

	// 学習の優先度を考えてここは後で実装する

	// cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("REGION")))
	// if err != nil {
	// }

	// client := sesv2.NewFromConfig(cfg)

	// input := &sesv2.SendEmailInput{
	// 	FromEmailAddress: ,
	// }

	return nil
}
