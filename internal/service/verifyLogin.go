package service

type LoginInfo struct {
	MailAddress string `json:"mail_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

func VerifyLogin(loginInfo LoginInfo) {

}
