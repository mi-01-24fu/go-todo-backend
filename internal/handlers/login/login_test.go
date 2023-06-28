package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/login"
	service "github.com/mi-01-24fu/go-todo-backend/internal/service/login"
)

func TestService_LoginHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userInfo := login.UserInfo{
		MailAddress: "inogan38@gmail.com",
		UserName:    "mifu",
	}
	verifyLoginResult := login.VerifyLoginResult{
		UserID:    1,
		LoginFlag: true,
	}

	inputJSON, _ := json.Marshal(&userInfo)
	reqBody := bytes.NewBufferString(string(inputJSON))

	serviceRepoMock := service.NewMockLoginRepository(ctrl)
	serviceRepoMock.EXPECT().VerifyLogin(gomock.Any(), userInfo).Return(verifyLoginResult, nil)

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		s       Service
		args    args
		want    login.VerifyLoginResult
		wantErr bool
	}{
		{
			"正常/VerifyLoginResultとnilを返却する",
			Service{Repo: serviceRepoMock},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", reqBody),
			},
			verifyLoginResult,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.LoginHandler(tt.args.w, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.LoginHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.LoginHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
