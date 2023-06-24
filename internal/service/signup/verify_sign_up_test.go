package signup

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	mockSignUp "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
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

	verifySignUpResult := signup.VerifySignUpResult{
		UserID:    1,
		LoginFlag: true,
	}

	type args struct {
		signupInfo signup.NewMemberInfo
	}
	tests := []struct {
		name string
		// setup関数の型定義だけする
		// そうすることで渡される引数は統一されるものの、期待値と戻り値を
		// 各テストケースごとにカスタマイズすることができる
		setup       func(*mockSignUp.MockAccessSignUp)
		args        args
		want        signup.VerifySignUpResult
		wantErr     bool
		errorString error
	}{
		{
			"正常/NewMemberInfoとnilを返却する",
			// ここではなにも実行されていない
			// 呼び出されたときにどういった挙動をするかの設定をしているだけ
			// 渡された引数の型MockAccessSignUpはインターフェースを満たしている
			func(msu *mockSignUp.MockAccessSignUp) {
				msu.EXPECT().SignUp(conversionReqBody("mifu", "inogan38@gmail.com")).Return(verifySignUpResult, nil)
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
			nil,
			args{
				conversionReqBody("", "inogan38@gmail.com"),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("UserName が入力されていません"),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/UserNameが2文字",
			nil,
			args{
				conversionReqBody("mi", "inogan38@gmail.com"),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("UserName の文字数が不正です"),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/UserNameが14文字",
			nil,
			args{
				conversionReqBody("14111111111111", "inogan38@gmail.com"),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("UserName の文字数が不正です"),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressが空",
			nil,
			args{
				conversionReqBody("mifu", ""),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("MailAddress が入力されていません"),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressの形式が不正「@」",
			nil,
			args{
				conversionReqBody("mifu", "inogan38gmail.com"),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("MailAddress の形式が正しくありません"),
		},
		{
			"準正常/NewMemberInfo{}とerrorを返却する/MailAddressの形式が不正「.」",
			nil,
			args{
				conversionReqBody("mifu", "inogan38@gmailcom"),
			},
			signup.VerifySignUpResult{},
			true,
			errors.New("MailAddress の形式が正しくありません"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// mocksignUpRepoにはMockAccessSignUp構造体が入る
			// この構造体はSignUpメソッドを定義しているインターフェースを満たす
			// 構造体。つまりこの構造体をもとにSignUpを呼び出すと、
			// Mock版のSignUpメソッドを呼び出せる
			// そうすることで実処理とテストで呼び出し先部分を変更することができる
			mocksignUpRepo := mockSignUp.NewMockAccessSignUp(ctrl)
			if tt.setup != nil {
				// ここで実際に各テストケースで記載しているsetup関数を呼出している
				// 渡している引数は構造体
				// そしてsetup関数内でテスト対象コードが渡すべき値や呼び出し先の設定をしている
				tt.setup(mocksignUpRepo)
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

func conversionReqBody(userName, mailAddress string) signup.NewMemberInfo {
	userInfo := signup.NewMemberInfo{
		UserName:    userName,
		MailAddress: mailAddress,
	}
	return userInfo
}