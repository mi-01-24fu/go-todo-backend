package mail_address

import (
	"regexp"
)

const validMailAddress = true
const invalidMailAddress = false
const regex = `^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`

// IsEmpty は mailAddress が空文字かの結果を返します
func IsEmpty(mailAddress string) bool {
	if mailAddress == "" {
		return invalidMailAddress
	}
	return validMailAddress
}

// CheckValidation は mailAddress が正しい形式化の結果を返します
func CheckValidation(mailAddress string) bool {
	if !regexp.MustCompile(regex).MatchString(mailAddress) {
		return invalidMailAddress
	}
	return validMailAddress
}
