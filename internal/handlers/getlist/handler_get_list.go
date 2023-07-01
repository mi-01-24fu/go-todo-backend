package getlist

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/getlist"
	getList "github.com/mi-01-24fu/go-todo-backend/internal/service/getlist"
)

// TODOGetHandler は VerifyGetTODOList を保持する構造体
type TODOGetHandler struct {
	GetTODORepo getList.VerifyGetTODOList
}

// NewGetListHandler は GetListHandler を生成して返却するコンストラクタ関数
func NewGetListHandler(g getList.VerifyGetTODOList) *TODOGetHandler {
	return &TODOGetHandler{GetTODORepo: g}
}

// GetTODOList は TODOList を取得するためのハンドラ関数
func (g TODOGetHandler) GetTODOList(w http.ResponseWriter, req *http.Request) {

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
