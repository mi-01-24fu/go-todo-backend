package addition

import (
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/service/addition"
)

// ServiceAddition は データ追加のハンドリングを行うインターフェース
type ServiceAddition interface {
	TaskAddition(w http.ResponseWriter, req *http.Request)
}

// TaskAdditionImpl は VerifyAdditionImpl を内部に保持する構造体
type TaskAdditionImpl struct {
	ServiceRepo *addition.VerifyAdditionImpl
}

// NewAdditionImple は新しい AdditionImple を作成し返却する
func NewAdditionImple(v *addition.VerifyAdditionImpl) *TaskAdditionImpl {
	return &TaskAdditionImpl{ServiceRepo: v}
}

// TaskAddition は データ追加のハンドリングを行う関数
func (a TaskAdditionImpl) TaskAddition(w http.ResponseWriter, req *http.Request) {}
