package configuration

import (
	"errors"
	"os"
	"strconv"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
)

// UseSES は AWS SES を利用するかの判定結果を返却します
func UseSES() (bool, error) {
	strFlag := os.Getenv("USESESFLAG")
	boolFlag, err := strconv.ParseBool(strFlag)
	if err != nil {
		return false, errors.New(consts.SystemError)
	}
	return boolFlag, nil
}
