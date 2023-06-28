package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/signup"
	mock "github.com/mi-01-24fu/go-todo-backend/internal/service/signup"
)

func TestSignUpService_SignUp(t *testing.T) {

	inputJSON, _ := json.Marshal("")
	reqBody := bytes.NewBufferString(string(inputJSON))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}

	tests := []struct {
		name    string
		setup   func(*mock.MockSignUp)
		args    args
		want    signup.VerifySignUpResult
		wantErr bool
	}{
		{
			"正常/VerifySignUpResultとnilを返却する",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(signup.RegistrationRequest{UserName: "mifu", MailAddress: "inogan38@gmail.com"}).Return(signup.VerifySignUpResult{UserID: 1, LoginFlag: true}, nil)
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("mifu", "inogan38@gmail.com")),
			},
			signup.VerifySignUpResult{
				UserID:    1,
				LoginFlag: true,
			},
			false,
		},
		{
			"準正常/VerifySignUpResult{}とerrorを返却する/UserNameが空",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(signup.RegistrationRequest{UserName: "", MailAddress: "inogan38@gmail.com"}).Return(signup.VerifySignUpResult{}, errors.New(consts.EmptyUserName))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("", "inogan38@gmail.com")),
			},
			signup.VerifySignUpResult{
				UserID:    0,
				LoginFlag: false,
			},
			true,
		},
		{
			"準正常/VerifySignUpResult{}とerrorを返却する/MailAddressが空",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(signup.RegistrationRequest{UserName: "mifu", MailAddress: ""}).Return(signup.VerifySignUpResult{}, errors.New(consts.EmptyMailAddress))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("mifu", "")),
			},
			signup.VerifySignUpResult{
				UserID:    0,
				LoginFlag: false,
			},
			true,
		},
		{
			"異常/VerifySignUpResult{}とerrorを返却する/RequestBodyがnil",
			nil,
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", nil),
			},
			signup.VerifySignUpResult{
				UserID:    0,
				LoginFlag: false,
			},
			true,
		},
		{
			"異常/VerifySignUpResult{}とerrorを返却する/RequestBodyがnil",
			nil,
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", reqBody),
			},
			signup.VerifySignUpResult{
				UserID:    0,
				LoginFlag: false,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			signUpRepoMock := mock.NewMockSignUp(ctrl)
			if tt.setup != nil {
				tt.setup(signUpRepoMock)
			}

			s := SignUpService{SignUpRepo: signUpRepoMock}
			got, err := s.SignUp(tt.args.w, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUpService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUpService.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func conversionReqBody(userName, mailAddress string) *bytes.Buffer {
	userInfo := signup.RegistrationRequest{
		UserName:    userName,
		MailAddress: mailAddress,
	}
	inputJSON, _ := json.Marshal(&userInfo)
	reqBody := bytes.NewBufferString(string(inputJSON))
	return reqBody
}
