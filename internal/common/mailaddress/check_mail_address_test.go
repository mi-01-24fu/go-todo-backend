package mailaddress

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		mailAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常/trueを返却する",
			args{"inogan38@gmail.com"},
			true,
		},
		{
			"異常/falseを返却する",
			args{""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.mailAddress); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckValidation(t *testing.T) {
	type args struct {
		mailAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常/trueを返却する/.com",
			args{mailAddress: "inogan38@gmail.com"},
			true,
		},
		{
			"正常/trueを返却/.jp",
			args{mailAddress: "inogan38@gmail.jp"},
			true,
		},
		{
			"正常/trueを返却/@icloud",
			args{mailAddress: "inogan38@icloud.com"},
			true,
		},
		{
			"異常/falseを返却/@無し",
			args{mailAddress: "inogan38icloud.com"},
			false,
		},
		{
			"異常/falseを返却/.無し",
			args{mailAddress: "inogan38@icloudcom"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidation(tt.args.mailAddress); got != tt.want {
				t.Errorf("CheckValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
