package verifySignup

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
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
	mock "github.com/mi-01-24fu/go-todo-backend/internal/service/verifySignup"
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
		want    verifySignup.VerifySignUpResponse
		wantErr bool
	}{
		{
			"正常/VerifySignUpResultとnilを返却する",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(verifySignup.VerifySignUpRequest{UserName: "mifu", MailAddress: "inogan38@gmail.com"}).Return(verifySignup.VerifySignUpResponse{UserID: 1, LoginFlag: true}, nil)
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("mifu", "inogan38@gmail.com")),
			},
			verifySignup.VerifySignUpResponse{
				UserID:    1,
				LoginFlag: true,
			},
			false,
		},
		{
			"準正常/VerifySignUpResult{}とerrorを返却する/UserNameが空",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(verifySignup.VerifySignUpRequest{UserName: "", MailAddress: "inogan38@gmail.com"}).Return(verifySignup.VerifySignUpResponse{}, errors.New(consts.EmptyUserName))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("", "inogan38@gmail.com")),
			},
			verifySignup.VerifySignUpResponse{
				UserID:    0,
				LoginFlag: false,
			},
			true,
		},
		{
			"準正常/VerifySignUpResult{}とerrorを返却する/MailAddressが空",
			func(msu *mock.MockSignUp) {
				msu.EXPECT().VerifySignUp(verifySignup.VerifySignUpRequest{UserName: "mifu", MailAddress: ""}).Return(verifySignup.VerifySignUpResponse{}, errors.New(consts.EmptyMailAddress))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/signup", conversionReqBody("mifu", "")),
			},
			verifySignup.VerifySignUpResponse{
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
			verifySignup.VerifySignUpResponse{
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
			verifySignup.VerifySignUpResponse{
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
	userInfo := verifySignup.VerifySignUpRequest{
		UserName:    userName,
		MailAddress: mailAddress,
	}
	inputJSON, _ := json.Marshal(&userInfo)
	reqBody := bytes.NewBufferString(string(inputJSON))
	return reqBody
}
