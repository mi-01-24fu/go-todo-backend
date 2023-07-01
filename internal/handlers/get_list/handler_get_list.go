package get_list

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
	getList "github.com/mi-01-24fu/go-todo-backend/internal/service/get_list"
)

// GetListHandler は VerifyGetTODOList を保持する構造体
type GetListHandler struct {
	GetTODORepo getList.VerifyGetTODOList
}

// NewGetListHandler は GetListHandler を生成して返却するコンストラクタ関数
func NewGetListHandler(g getList.VerifyGetTODOList) *GetListHandler {
	return &GetListHandler{GetTODORepo: g}
}

// GetTODOList は TODOList を取得するためのハンドラ関数
func (g GetListHandler) GetTODOList(w http.ResponseWriter, req *http.Request) {

	// リクエストデータが読み取れるか確認
	todoRequest, err := checkGetTODOInput(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODOList 取得処理呼出し
	result, err := g.GetTODORepo.GetTODOList(todoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンス返却
	fmt.Println(result)
	createResponse(w, result)
}

// checkGetTODOInput は望むリクエストデータが送られてきているかを確認します
func checkGetTODOInput(req *http.Request) (access.GetTODORequest, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return access.GetTODORequest{}, errors.New(consts.BadInput)
	}

	var todoRequest access.GetTODORequest

	if err := json.Unmarshal(body, &todoRequest); err != nil {
		return access.GetTODORequest{}, errors.New(consts.BadInput)
	}
	return todoRequest, nil
}

// createResponse はレスポンスデータを加工して返却する
func createResponse(w http.ResponseWriter, todoList getList.ResponseList) {
	res, err := json.Marshal(todoList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
