package user_name

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常/trueを返却する",
			args{"mifu"},
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
			if got := IsEmpty(tt.args.userName); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckLength(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常/trueを返却する/3文字",
			args{"333"},
			true,
		},
		{
			"正常/trueを返却する/12文字",
			args{"121111111111"},
			true,
		},
		{
			"異常/falseを返却する/2文字",
			args{"22"},
			false,
		},
		{
			"異常/falseを返却する/13文字",
			args{"1311111111111"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLength(tt.args.userName); got != tt.want {
				t.Errorf("CheckLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
