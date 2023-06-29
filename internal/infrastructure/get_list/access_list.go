package get_list

import (
	"database/sql"
	"fmt"
)

// AccessTODO は TODOSテーブルへアクセスするためのインターフェース
type AccessTODO interface {
	CheckID(GetTODORequest) (bool, error)
	GetTODO(GetTODORequest) (GetTODOList, error)
}

// AccessTODOImpl は AccessTODO を実装する構造体
type AccessTODOImpl struct {
	DB *sql.DB
}

// GetTODORequest は クライアントから渡された ID を格納する構造体
type GetTODORequest struct {
	ID int
}

// GetTODOList は TODOS テーブルから取得したTODOList を格納する構造体
type GetTODOList struct {
	ID          int
	ActiveTask  string
	Description string
}

// NewAccessTODOImpl は AccessTODOImpl を生成して返却するコンストラクタ関数
func NewAccessTODOImpl(db *sql.DB) *AccessTODOImpl {
	return &AccessTODOImpl{DB: db}
}

// VerifyID は 受け取ったIDが TODOS テーブルに存在するかの結果を返却する
func (a AccessTODOImpl) CheckID(id GetTODORequest) (bool, error) {
	fmt.Println("11")
	return false, nil
}

// GetTODOList は TODOS テーブルから ID に紐づく TODOList の一覧を返却する
func (a AccessTODOImpl) GetTODO(id GetTODORequest) (GetTODOList, error) {
	fmt.Println("11")

	return GetTODOList{}, nil
}
