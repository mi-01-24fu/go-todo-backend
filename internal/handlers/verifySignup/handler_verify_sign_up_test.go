package verifySignup

import (
	"net/http"
	"net/smtp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/service/verifySignup"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
)

func TestHandlerSignUpRepo_VerifySignUp(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockPreparationSingUp := NewMockPreparationSingUp(ctrl)
	mockHandlerSignUpRepo := NewHandlerSignUpRepo(mockPreparationSingUp)

	erifySignUpResponse := access.VerifySignUpRequest{
		UserName: "mifu",
		MailAddress: "inogan38@gmail.com",
	}

	verifySignUpResponse := access.VerifySignUpResponse {
		AuthenticationNumber: 1234,
	}


	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name  string
		setup func(*verifySignup.MockPreparationSingUp)
		h     *HandlerSignUpRepo
		args  args
	}{
		{
			"正常",
			func(mpsu *verifySignup.MockPreparationSingUp) {
				mpsu.EXPECT().VerifySignUp().Return(),
			},
			mockHandlerSignUpRepo,
			args {

			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.VerifySignUp(tt.args.w, tt.args.req)
		})
	}
}
