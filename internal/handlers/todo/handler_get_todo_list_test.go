package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/todo"
	"github.com/mi-01-24fu/go-todo-backend/internal/service/todo"
	"github.com/stretchr/testify/assert"
)

func TestTODOListHandler_GetTODOList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getTODORequest := access.GetTODORequest{ID: 1}
	getTODOList := access.GetTODOList{ID: 1, ActiveTask: "sample", Description: "sample"}

	inputJSON, _ := json.Marshal("")
	errorReqBody := bytes.NewBufferString(string(inputJSON))

	url := "http://localhost:8080/gettodo"

	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name        string
		setup       func(*todo.MockVerifyGetTODOList)
		args        args
		want        access.GetTODOList
		wantErr     bool
		errorString string
	}{
		{
			"正常/TODOリストとnilを返却する",
			func(mus *todo.MockVerifyGetTODOList) {
				mus.EXPECT().GetTODOList(getTODORequest).Return(getTODOList, nil)
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodGet, url, conversionReqBody(1)),
			},
			getTODOList,
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
			access.GetTODOList{},
			true,
			consts.BadInput,
		},
		{
			"異常/GetTODOListの戻り値エラー",
			func(mus *todo.MockVerifyGetTODOList) {
				mus.EXPECT().GetTODOList(getTODORequest).Return(access.GetTODOList{}, errors.New(consts.SystemError))
			},
			args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodGet, url, conversionReqBody(1)),
			},
			access.GetTODOList{},
			true,
			consts.SystemError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockVerifyGetTODOList := todo.NewMockVerifyGetTODOList(ctrl)

			if tt.setup != nil {
				tt.setup(mockVerifyGetTODOList)
			}
			g := TODOListHandler{GetTODORepo: mockVerifyGetTODOList}

			g.GetTODOList(tt.args.w, tt.args.req)

			fmt.Println(tt.args.w)

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

func conversionResBody(res access.GetTODOList) *bytes.Buffer {
	responseJSON, _ := json.Marshal(&res)
	resBody := bytes.NewBufferString(string(responseJSON))
	return resBody
}
