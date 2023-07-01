package getlist

import (
	"errors"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/getlist"
)

// VerifyGetTODOList は TODOListを取得するインターフェース
type VerifyGetTODOList interface {
	GetTODOList(access.GetTODORequest) (ResponseList, error)
}

// GetService は VerifyGetTODOList を満たす構造体
type GetService struct {
	AccessTODORepo access.AccessTODO
}

// NewGetService は GetService を生成して返却するコンストラクタ関数
func NewGetService(a access.AccessTODO) VerifyGetTODOList {
	return &GetService{AccessTODORepo: a}
}

// ResponseList は クライアントへ返却するデータを格納する構造体
type ResponseList struct {
	Lists    access.GetLists
	ListFlag bool
}

// GetTODOList は TODOList を取得する
func (g GetService) GetTODOList(req access.GetTODORequest) (ResponseList, error) {

	err := checkValidation(req)
	if err != nil {
		return ResponseList{}, err
	}

	emptyRecord, err := g.AccessTODORepo.CheckID(req)
	if err != nil {
		return ResponseList{}, err
	}

	if !emptyRecord {
		return ResponseList{nil, false}, nil
	}

	getList, err := g.AccessTODORepo.GetTODOList(req)
	if err != nil {
		return ResponseList{}, err
	}

	responseData := createResponse(getList)
	return responseData, nil
}

// checkValidation はリクエストデータのバリデーションチェックを行います
func checkValidation(data access.GetTODORequest) error {
	if data.ID == 0 {
		return errors.New(consts.SystemError)
	}
	return nil
}

// createResponse は クライアントへ返却するレスポンスデータを作成する
func createResponse(data access.GetLists) ResponseList {
	if data == nil {
		return ResponseList{nil, false}
	}
	return ResponseList{data, true}
}
