package addition

import "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/addition"

// ReqestAddition はクライアントから渡されたパラメーターを格納する構造体
type ReqestAddition struct{}

// ResponseAddition はクライアントへ返却するデータを格納する構造体
type ResponseAddition struct{}

// VerifyAddition は DB へデータを登録するためのインターフェース
type VerifyAddition interface {
	Insert()
}

// VerifyAdditionImpl は VerifyAddition を満たす構造体
type VerifyAdditionImpl struct {
	AccessRepo *addition.TaskAdditionImpl
}

// NewVerifyAdditionImpl は 新しい VerifyAdditionImpl を作成し返却する
func NewVerifyAdditionImpl(a *addition.TaskAdditionImpl) *VerifyAdditionImpl {
	return &VerifyAdditionImpl{
		AccessRepo: a,
	}
}

// Insert は TODO データを DB へ登録する
func (v VerifyAdditionImpl) Insert() {}
