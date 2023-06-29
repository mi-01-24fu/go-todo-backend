package get_list

import (
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
)

// VerifyGetTODOList は TODOListを取得するインターフェース
type VerifyGetTODOList interface {
	GetTODOList(access.GetTODORequest) (access.GetTODOList, error)
}

// GetService は VerifyGetTODOList を満たす構造体
type GetService struct {
	AccessTODORepo *access.AccessTODOImpl
}

// NewGetService は GetService を生成して返却するコンストラクタ関数
func NewGetService(a *access.AccessTODOImpl) *GetService {
	return &GetService{AccessTODORepo: a}
}

// GetTODOList は TODOList を取得する
func (g GetService) GetTODOList(req access.GetTODORequest) (access.GetTODOList, error) {
	return g.AccessTODORepo.GetTODO(req)
}
