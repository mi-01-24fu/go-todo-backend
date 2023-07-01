package addition

import "database/sql"

// TaskAddition は TODO データを登録するためのインターフェース
type TaskAddition interface {
	ADD()
}

// TaskAdditionImpl は DB へ接続するための情報を格納する構造体
type TaskAdditionImpl struct {
	DB *sql.DB
}

// NewAdditionTaskImpl は新しい AdditionTaskImpl を作成し返却する
func NewAdditionTaskImpl(db *sql.DB) *TaskAdditionImpl {
	return &TaskAdditionImpl{
		DB: db,
	}
}

// ADD は TODO データを登録する
func (a TaskAdditionImpl) ADD() {}
