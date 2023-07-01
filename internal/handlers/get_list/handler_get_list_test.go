package get_list

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
	"github.com/mi-01-24fu/go-todo-backend/internal/service/get_list"
	getList "github.com/mi-01-24fu/go-todo-backend/internal/service/get_list"
	"github.com/stretchr/testify/assert"
)

func TestTODOListHandler_GetTODOList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getTODORequest := access.GetTODORequest{ID: 1}
	//getTODOList := access.GetTODOList{ID: 1, ActiveTask: "sample", Description: "sample"}

	inputJSON, _ := json.Marshal("")
	errorReqBody := bytes.NewBufferString(string(inputJSON))

	getLists := access.GetLists{}
	todoList := access.GetTODOList{1, 1, "activeTask", "Description"}
	getLists = append(getLists, todoList)
	responseList := getList.ResponseList{getLists, true}

	url := "http://localhost:8080/gettodo"

	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name        string
		setup       func(*getList.MockVerifyGetTODOList)
		args        args
		want        getList.ResponseList
		wantErr     bool
		errorString string
	}{
		{
			"正常/TODOリストとnilを返却する",
			func(mus *getList.MockVerifyGetTODOList) {
				mus.EXPECT().GetTODOList(getTODORequest).Return(responseList, nil)
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodGet, url, conversionReqBody(1)),
			},
			responseList,
			false,
			"",
		},
		{
			"異常/リクエストデータの読み取りエラー",
			nil,
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodGet, url, errorReqBody),
			},
			getList.ResponseList{},
			true,
			consts.BadInput,
		},
		{
			"異常/GetTODOListの戻り値エラー",
			func(mus *get_list.MockVerifyGetTODOList) {
				mus.EXPECT().GetTODOList(getTODORequest).Return(getList.ResponseList{}, errors.New(consts.SystemError))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodGet, url, conversionReqBody(1)),
			},
			getList.ResponseList{},
			true,
			consts.SystemError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockVerifyGetTODOList := get_list.NewMockVerifyGetTODOList(ctrl)

			if tt.setup != nil {
				tt.setup(mockVerifyGetTODOList)
			}
			g := GetListHandler{GetTODORepo: mockVerifyGetTODOList}

			g.GetTODOList(tt.args.w, tt.args.req)

			if !tt.wantErr {
				assert.Equal(t, http.StatusOK, tt.args.w.Code)
				assert.Equal(t, "application/json", tt.args.w.Header().Get("Content-Type"))
				assert.Equal(t, conversionResBody(tt.want).String(), tt.args.w.Body.String())
			}

			if tt.wantErr {
				assert.NotEqual(t, http.StatusOK, tt.args.w.Code)
				actualErrorMessage := strings.TrimSpace(tt.args.w.Body.String()) // strings.TrimSpace()を使用して改行を除去
				assert.Equal(t, tt.errorString, actualErrorMessage)
			}
		})
	}
}

func conversionReqBody(id int) *bytes.Buffer {
	todoInfo := access.GetTODORequest{ID: id}
	inputJSON, _ := json.Marshal(&todoInfo)
	reqBody := bytes.NewBufferString(string(inputJSON))
	return reqBody
}

func conversionResBody(res getList.ResponseList) *bytes.Buffer {
	responseJSON, _ := json.Marshal(&res)
	resBody := bytes.NewBufferString(string(responseJSON))
	return resBody
}
