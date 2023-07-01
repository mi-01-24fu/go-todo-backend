package get_list

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mi-01-24fu/go-todo-backend/internal/consts"
	access "github.com/mi-01-24fu/go-todo-backend/internal/infrastructure/get_list"
)

func TestGetService_GetTODOList(t *testing.T) {

	ctrl := gomock.NewController(t)
	getLists := access.GetLists{}
	todoList := access.GetTODOList{1, 1, "activeTask", "Description"}
	getLists = append(getLists, todoList)
	responseList := ResponseList{getLists, true}

	request := access.GetTODORequest{ID: 1}

	type setup struct {
		checkID     func(*access.MockAccessTODO)
		getTODOList func(*access.MockAccessTODO)
	}
	type args struct {
		req access.GetTODORequest
	}
	type want struct {
		list ResponseList
	}
	tests := []struct {
		name        string
		setup       setup
		args        args
		want        want
		errorString error
		wantErr     bool
	}{
		{
			"正常/TODOリストとnilを返却する",
			setup{
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().CheckID(request).Return(true, nil)
				},
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().GetTODOList(request).Return(getLists, nil)
				},
			},
			args{
				request,
			},
			want{
				responseList,
			},
			nil,
			false,
		},
		{
			"正常/TODOリストの空(listFlag false)とnilを返却する/登録しているTODOリストが無い",
			setup{
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().CheckID(request).Return(false, nil)
				},
				nil,
			},
			args{
				request,
			},
			want{
				ResponseList{nil, false},
			},
			nil,
			false,
		},
		{
			"異常/TODOリストの空データとerrorを返却する/CheckId_DB接続時エラー",
			setup{
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().CheckID(request).Return(false, errors.New(consts.SystemError))
				},
				nil,
			},
			args{
				request,
			},
			want{
				ResponseList{nil, false},
			},
			errors.New(consts.SystemError),
			true,
		},
		{
			"異常/TODOリストの空データとerrorを返却する/GetTODOList_DB接続時エラー",
			setup{
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().CheckID(request).Return(true, nil)
				},
				func(mus *access.MockAccessTODO) {
					mus.EXPECT().GetTODOList(request).Return(nil, errors.New(consts.SystemError))
				},
			},
			args{
				request,
			},
			want{
				ResponseList{nil, false},
			},
			errors.New(consts.SystemError),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock := access.NewMockAccessTODO(ctrl)
			if tt.setup.checkID != nil {
				tt.setup.checkID(mock)
			}

			if tt.setup.getTODOList != nil {
				tt.setup.getTODOList(mock)
			}

			mockRepo := NewGetService(mock)

			got, err := mockRepo.GetTODOList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetService.GetTODOList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				if err.Error() != tt.errorString.Error() {
					t.Errorf("GetService.GetTODOList() error = %v, wantErr %v", err, tt.errorString)
				}
			}

			if !reflect.DeepEqual(got, tt.want.list) {
				t.Errorf("GetService.GetTODOList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// createResponseData は クライアントへ返却するデータを生成します
func createResponseData(data access.GetTODOList, listFlag bool) ResponseList {
	getLists := access.GetLists{}
	todoList := access.GetTODOList{data.ID, data.UserID, data.ActiveTask, data.Description}
	getLists = append(getLists, todoList)
	return ResponseList{getLists, listFlag}
}
