package addition

import "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/addition"

type AdditionReqest struct{}
type AdditionResponse struct{}

type VerifyAddition interface {
	Insert()
}

type VerifyAdditionImpl struct {
	AccessRepo *addition.AdditionTaskImpl
}

func NewVerifyAdditionImpl(a *addition.AdditionTaskImpl) *VerifyAdditionImpl {
	return &VerifyAdditionImpl{
		AccessRepo: a,
	}
}

func (v VerifyAdditionImpl) Insert() {}
