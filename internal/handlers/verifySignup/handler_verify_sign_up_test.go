package verifySignup

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
	"github.com/mi-01-24fu/go-todo-backend/internal/service/verifySignup"
	"github.com/stretchr/testify/assert"
)

func TestHandlerSignUpRepo_VerifySignUp(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockPreparationSingUp := verifySignup.NewMockPreparationSingUp(ctrl)
	mockHandlerSignUpRepo := NewHandlerSignUpRepo(mockPreparationSingUp)

	verifySignUpRequest := access.VerifySignUpRequest{
		UserName:    "mifu",
		MailAddress: "inogan38@gmail.com",
	}

	goodVerifySignUpResponse := access.VerifySignUpResponse{
		AuthenticationNumber: 1234,
	}

	badVerifySignUpResponse := access.VerifySignUpResponse{
		AuthenticationNumber: 123,
	}

	url := "http://localhost:8080/verify-signup"

	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name      string
		setup     func(*verifySignup.MockPreparationSingUp)
		h         *HandlerSignUpRepo
		args      args
		want      access.VerifySignUpResponse
		wantErr   bool
		errString string
	}{
		{
			"正常/4桁の乱数を返却する",
			func(mpsu *verifySignup.MockPreparationSingUp) {
				mpsu.EXPECT().VerifySignUp(context.Background(), verifySignUpRequest).Return(goodVerifySignUpResponse, nil)
			},
			mockHandlerSignUpRepo,
			args{
				// http.ResponseWriter はレスポンスを返却するためのメソッドを
				// 提供するインターフェース
				// そして httptest.NewRecorder() は http.ResponseWriter インターフェースを
				// 実装している ResponseRecorder 構造体を返却する
				w: httptest.NewRecorder(),
				// httptest.NewRequest は疑似的な HTTP リクエストを
				// 生成できる
				req: httptest.NewRequest(http.MethodPost, url, conversionReqBody("mifu", "inogan38@gmail.com")),
			},
			goodVerifySignUpResponse,
			false,
			"",
		},
		{
			"異常/VerifySignUpが4桁以外の数字を返却",
			func(mpsu *verifySignup.MockPreparationSingUp) {
				mpsu.EXPECT().VerifySignUp(context.Background(), verifySignUpRequest).Return(badVerifySignUpResponse, nil)
			},
			mockHandlerSignUpRepo,
			args{
				// http.ResponseWriter はレスポンスを返却するためのメソッドを
				// 提供するインターフェース
				// そして httptest.NewRecorder() は http.ResponseWriter インターフェースを
				// 実装している ResponseRecorder 構造体を返却する
				w: httptest.NewRecorder(),
				// httptest.NewRequest は疑似的な HTTP リクエストを
				// 生成できる
				req: httptest.NewRequest(http.MethodPost, url, conversionReqBody("mifu", "inogan38@gmail.com")),
			},
			badVerifySignUpResponse,
			true,
			consts.SystemError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(mockPreparationSingUp)

			tt.h.VerifySignUp(tt.args.w, tt.args.req)

			if !tt.wantErr {
				assert.Equal(t, http.StatusOK, tt.args.w.Code)
				assert.Equal(t, "application/json", tt.args.w.Header().Get("Content-Type"))
				assert.Equal(t, conversionResBody(t, tt.want).String(), tt.args.w.Body.String())
			}

			if tt.wantErr {
				assert.NotEqual(t, http.StatusOK, tt.args.w.Code)
				actualErrorMessage := strings.TrimSpace(tt.args.w.Body.String()) // strings.TrimSpace()を使用して改行を除去
				assert.Equal(t, tt.errString, actualErrorMessage)
			}
		})
	}
}

// conversionReqBody はリクエストBodyを生成する
func conversionReqBody(userName, mailAddress string) *bytes.Buffer {
	verifySignUpRequest := access.VerifySignUpRequest{
		UserName:    userName,
		MailAddress: mailAddress,
	}
	inputJSON, _ := json.Marshal(&verifySignUpRequest)
	reqBody := bytes.NewBufferString(string(inputJSON))
	return reqBody
}

// conversionResBody は期待するレスポンスを bytes.Buffer 型に変換する
func conversionResBody(t *testing.T, res access.VerifySignUpResponse) *bytes.Buffer {
	responseJSON, err := json.Marshal(&res)
	if err != nil {
		t.Errorf("conversionResBody() json marshal error =  %v", err)
	}
	return bytes.NewBufferString(string(responseJSON))
}
