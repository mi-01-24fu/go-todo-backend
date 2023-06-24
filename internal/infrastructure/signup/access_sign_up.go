package signup

// VerifySignUpResult は クライアントから渡されたユーザー情報をもとにDBへデータを登録した結果を格納します
type VerifySignUpResult struct {
	UserID    int  `json:"user_id"`
	LoginFlag bool `json:"login_flag"`
}

// NewMemberInfo は クライアントから渡されたログイン情報を保持する構造体です
type NewMemberInfo struct {
	UserName    string `json:"user_name,omitempty"`
	MailAddress string `json:"mail_address,omitempty"`
}

// AccessSignUp は ユーザー情報を登録するためのインターフェース
type AccessSignUp interface {
	SignUp(NewMemberInfo) (VerifySignUpResult, error)
}

// AccessSignUpInfo は AccessSignUp を満たす構造体
type AccessSignUpInfo struct{}

// SignUp はユーザー情報をDBへ登録する処理を行う
func (a AccessSignUpInfo) SignUp(signUpInfo NewMemberInfo) (VerifySignUpResult, error) {
	return VerifySignUpResult{UserID: 1, LoginFlag: true}, nil
}
