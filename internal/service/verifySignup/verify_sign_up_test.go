package verifySignup

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
	mockSignUp "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
	"github.com/stretchr/testify/assert"
)

func TestSignUpInfo_VerifySignUp(t *testing.T) {

	// gomockのコントローラーを作成
	// ctrlがモックの期待値が満たされたかどうかのチェックなどを行う
	ctrl := gomock.NewController(t)

	// これはテストが終了するとき（関数がリターンするとき）、
	// gomockのコントローラーに対してFinishメソッドが
	// 呼び出されるように設定するもので、これによってgomockは
	// 全ての期待された呼び出しが行われたかどうかをチェックする
	// もし期待された呼び出しが全て行われていなければ
	// テストは失敗する
	defer ctrl.Finish()

	verifySignUpResult := verifySignup.VerifySignUpResponse{
		UserID:    1,
		LoginFlag: true,
	}

	type setup struct {
		signUp func(*mockSignUp.MockService)
		count  func(*mockSignUp.MockService)
	}

	type args struct {
		signupInfo verifySignup.VerifySignUpRequest
	}
	tests := []struct {
		name string
		// setup関数の型定義だけする
		// そうすることで渡される引数は統一されるものの、期待値と戻り値を
		// 各テストケースごとにカスタマイズすることができる
		setup       setup
		args        args
		want        verifySignup.VerifySignUpResponse
		wantErr     bool
		errorString error
	}{
		{
			"正常/NewMemberInfoとnilを返却する",
			// ここではなにも実行されていない
			// 呼び出されたときにどういった挙動をするかの設定をしているだけ
			// 渡された引数の型MockServiceはインターフェースを満たしている
			setup{
				func(msu *mockSignUp.MockService) {
					msu.EXPECT().SignUp(conversionReqBody("mifu", "inogan38@gmail.com")).Return(verifySignUpResult, nil)
				},
				func(msu *mockSignUp.MockService) {
					msu.EXPECT().Count("inogan38@gmail.com").Return(int64(0), nil)
				},
			},
			args{
				conversionReqBody("mifu", "inogan38@gmail.com"),
			},
			verifySignUpResult,
			false,
			nil,
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/UserNameが空",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("", "inogan38@gmail.com"),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.EmptyUserName),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/UserNameが2文字",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("mi", "inogan38@gmail.com"),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.InvalidUserNameLength),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/UserNameが14文字",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("14111111111111", "inogan38@gmail.com"),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.InvalidUserNameLength),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressが空",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("mifu", ""),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.EmptyMailAddress),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressの形式が不正「@」",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("mifu", "inogan38gmail.com"),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.NotMailAddressFormat),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressの形式が不正「.」",
			setup{
				nil,
				nil,
			},
			args{
				conversionReqBody("mifu", "inogan38@gmailcom"),
			},
			verifySignup.VerifySignUpResponse{},
			true,
			errors.New(consts.NotMailAddressFormat),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// mocksignUpRepoにはMockService構造体が入る
			// この構造体はSignUpメソッドを定義しているインターフェースを満たす
			// 構造体。つまりこの構造体をもとにSignUpを呼び出すと、
			// Mock版のSignUpメソッドを呼び出せる
			// そうすることで実処理とテストで呼び出し先部分を変更することができる
			mocksignUpRepo := mockSignUp.NewMockService(ctrl)
			if tt.setup.signUp != nil {
				// ここで実際に各テストケースで記載しているsetup関数を呼出している
				// 渡している引数は構造体
				// そしてsetup関数内でテスト対象コードが渡すべき値や呼び出し先の設定をしている
				tt.setup.signUp(mocksignUpRepo)
			}

			if tt.setup.count != nil {
				tt.setup.count(mocksignUpRepo)
			}
			// AccessRepoに設定しているのはmocksignUpRepo構造体
			// mocksignUpRepo構造体はSignUpのインターフェースを満たしているため、
			// こうすることでテスト対象側でSignUpを呼び出した際に、Mock版の
			// SignUpメソッドを呼び出すことができる
			signUpRepo := AccessInfo{AccessRepo: mocksignUpRepo}

			// signUpRepo.このように呼び出すことで、VerifySignUpに
			// signUpRepoを引数として渡さなくても、signUpRepo内の値が
			// 暗黙的にVerifySignUpに引き継がれる
			// つまりテスト対象コード側でs.AccessRepo.SignUp(signupInfo)
			// 上記のように呼び出しており、AccessRepoに格納した
			// 構造体がMockの構造体の為、そのMockの構造体に紐づくSignUp
			// を呼び出すことができ、Mock化することができる
			got, err := signUpRepo.VerifySignUp(tt.args.signupInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUpInfo.VerifySignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				assert.Equal(t, err, tt.errorString)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUpInfo.VerifySignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func conversionReqBody(userName, mailAddress string) verifySignup.VerifySignUpRequest {
	userInfo := verifySignup.VerifySignUpRequest{
		UserName:    userName,
		MailAddress: mailAddress,
	}
	return userInfo
}
