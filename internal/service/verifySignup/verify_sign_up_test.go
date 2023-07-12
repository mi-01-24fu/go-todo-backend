package verifySignup

import (
	"context"
	"reflect"
	"testing"

	"github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/verifySignup"
)

func TestPreparationSingUpImpl_VerifySignUp(t *testing.T) {
	type args struct {
		ctx         context.Context
		requestData verifySignup.VerifySignUpRequest
	}
	tests := []struct {
		name    string
		s       PreparationSingUpImpl
		args    args
		want    verifySignup.VerifySignUpResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.VerifySignUp(tt.args.ctx, tt.args.requestData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PreparationSingUpImpl.VerifySignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreparationSingUpImpl.VerifySignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
