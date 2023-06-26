package signup

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestServiceImpl_Count(t *testing.T) {
	type args struct {
		mailaddress string
	}
	tests := []struct {
		name    string
		a       ServiceImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Count(tt.args.mailaddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ServiceImpl.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_SignUp(t *testing.T) {
	type args struct {
		signUpInfo RegistrationRequest
	}
	tests := []struct {
		name    string
		a       ServiceImpl
		args    args
		want    VerifySignUpResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.SignUp(tt.args.signUpInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
