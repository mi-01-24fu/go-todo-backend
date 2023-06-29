package addition

import (
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/service/addition"
)

type AdditionService interface {
	AdditionTask(w http.ResponseWriter, req *http.Request)
}

type AdditionImple struct {
	ServiceRepo *addition.VerifyAdditionImpl
}

func NewAdditionImple(v *addition.VerifyAdditionImpl) *AdditionImple {
	return &AdditionImple{ServiceRepo: v}
}

func (a AdditionImple) AdditionTask(w http.ResponseWriter, req *http.Request) {

}
