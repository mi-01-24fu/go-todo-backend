package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	service "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
	"github.com/stretchr/testify/assert"
)

func TestService_LoginHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userInfo := login.RequestUser{
		MailAddress: "inogan38@gmail.com",
		UserName:    "mifu",
	}
	verifyLoginResult := login.ResponseUser{
		UserID:    1,
		LoginFlag: true,
	}

	inputJSON, _ := json.Marshal(&userInfo)
	reqBody := bytes.NewBufferString(string(inputJSON))

	errorInputJSON, _ := json.Marshal("")
	errorReqBody := bytes.NewBufferString(string(errorInputJSON))

	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name        string
		setup       func(*service.MockLoginRepository)
		args        args
		want        login.ResponseUser
		errorString string
		wantErr     bool
	}{
		{
			"正常/VerifyLoginResultとnilを返却する",
			func(mlr *service.MockLoginRepository) {
				mlr.EXPECT().VerifyLogin(gomock.Any(), userInfo).Return(verifyLoginResult, nil)
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", reqBody),
			},
			verifyLoginResult,
			"",
			false,
		},
		{
			"異常系/リクエストデータの読み取りエラー",
			nil,
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", errorReqBody),
			},
			login.ResponseUser{},
			consts.BadInput,
			true,
		},
		{
			"異常/VerifyLogin呼出しエラー",
			func(mlr *service.MockLoginRepository) {
				mlr.EXPECT().VerifyLogin(gomock.Any(), userInfo).Return(login.ResponseUser{}, errors.New(consts.SystemError))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", bytes.NewBufferString(string(inputJSON))),
			},
			login.ResponseUser{},
			consts.SystemError,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockVerifyLogin := service.NewMockLoginRepository(ctrl)

			if tt.setup != nil {
				tt.setup(mockVerifyLogin)
			}

			mocklogin := VerifyLoginHandler{LoginRepo: mockVerifyLogin}

			mocklogin.LoginHandler(tt.args.w, tt.args.req)

			if !tt.wantErr {
				assert.Equal(t, http.StatusOK, tt.args.w.Code)
				assert.Equal(t, "application/json", tt.args.w.Header().Get("Content-Type"))
				assert.Equal(t, conversionResBody(tt.want).String(), tt.args.w.Body.String())
			}

			if tt.wantErr {
				assert.NotEqual(t, http.StatusOK, tt.args.w.Code)
				actualErrorMessage := strings.TrimSpace(tt.args.w.Body.String()) // strings.TrimSpace()を使用して改行を除去
				assert.Equal(t, tt.errorString, actualErrorMessage)
			}
		})
	}
}

func conversionResBody(res login.ResponseUser) *bytes.Buffer {
	responseJSON, _ := json.Marshal(&res)
	resBody := bytes.NewBufferString(string(responseJSON))
	return resBody
}
