package username

const minNamelength = 3
const maxNamelength = 12
const validUserName = true
const invalidUserName = false

// IsEmpty は userName が空文字かの結果を返します
func IsEmpty(userName string) bool {
	if userName == "" {
		return invalidUserName
	}
	return validUserName
}

// CheckLength は userName の文字列長が正しいかの結果を返します
func CheckLength(userName string) bool {
	if len(userName) < minNamelength || len(userName) > maxNamelength {
		return invalidUserName
	}
	return validUserName
}
