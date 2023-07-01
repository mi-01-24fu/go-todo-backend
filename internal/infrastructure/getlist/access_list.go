package getlist

import (
	"context"
	"database/sql"
	"log"

	"github.com/mi-01-24fu/go-todo-backend/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AccessTODO は TODOSテーブルへアクセスするためのインターフェース
type AccessTODO interface {
	CheckID(GetTODORequest) (bool, error)
	GetTODOList(GetTODORequest) (GetLists, error)
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
	UserID      int
	ActiveTask  string
	Description string
}

// GetLists は DB から取得したデータをスライスで保持する
type GetLists []GetTODOList

// NewAccessTODOImpl は AccessTODOImpl を生成して返却するコンストラクタ関数
func NewAccessTODOImpl(db *sql.DB) AccessTODO {
	return &AccessTODOImpl{DB: db}
}

// CheckID は 受け取ったIDが TODOS テーブルに存在するかの結果を返却する
func (a AccessTODOImpl) CheckID(requestData GetTODORequest) (bool, error) {

	ctx := context.Background()

	record, err := models.Todos(
		qm.Where("user_id=?", requestData.ID),
	).Count(ctx, a.DB)

	if err != nil {
		log.Print(err)
		return false, err
	}
	if record == 0 {
		return false, nil
	}

	return true, nil
}

// GetTODOList は TODOS テーブルから ID に紐づく TODOList の一覧を返却する
func (a AccessTODOImpl) GetTODOList(requestData GetTODORequest) (GetLists, error) {

	ctx := context.Background()
	list := GetLists{}

	todoList, err := models.Todos(
		qm.Where("user_id=?", requestData.ID),
	).All(ctx, a.DB)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	for _, todo := range todoList {
		todoList := GetTODOList{
			todo.ID,
			todo.UserID.Int,
			todo.ActiveTask,
			todo.Description.String,
		}
		list = append(list, todoList)
	}

	return list, nil
}
