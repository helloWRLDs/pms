package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
)

func TestCreateTaskComment(t *testing.T) {
	tests := []struct {
		name     string
		creation *dto.TaskCommentCreation
		wantErr  bool
	}{
		{
			name: "create valid comment",
			creation: &dto.TaskCommentCreation{
				Body:   "This is a test comment",
				TaskId: "712b8a41-2351-4286-ad03-086eaee4c417",
				UserId: "f3cef382-559d-4248-9b02-9c0038725ab7",
			},
			wantErr: false,
		},
		{
			name: "create comment with empty body",
			creation: &dto.TaskCommentCreation{
				Body:   "",
				TaskId: "712b8a41-2351-4286-ad03-086eaee4c417",
				UserId: "f3cef382-559d-4248-9b02-9c0038725ab7",
			},
			wantErr: true,
		},
		{
			name: "create comment for non-existent task",
			creation: &dto.TaskCommentCreation{
				Body:   "This is a test comment",
				TaskId: "non-existent-task",
				UserId: "f3cef382-559d-4248-9b02-9c0038725ab7",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := logic.CreateTaskComment(context.Background(), tt.creation)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, comment)
			assert.Equal(t, tt.creation.Body, comment.Body)
			assert.Equal(t, tt.creation.TaskId, comment.TaskId)
			assert.Equal(t, tt.creation.UserId, comment.User.Id)
		})
	}
}

func TestListTaskComments(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.TaskCommentFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all comments",
			filter: &dto.TaskCommentFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list comments by task",
			filter: &dto.TaskCommentFilter{
				Page:    1,
				PerPage: 10,
				TaskId:  "712b8a41-2351-4286-ad03-086eaee4c417",
			},
			wantCount: 5,
			wantErr:   false,
		},
		{
			name: "list comments for non-existent task",
			filter: &dto.TaskCommentFilter{
				Page:    1,
				PerPage: 10,
				TaskId:  "non-existent-task",
			},
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := logic.ListTaskComments(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, comments)
			assert.Len(t, comments.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, int(comments.Page))
			assert.Equal(t, tt.filter.PerPage, int(comments.PerPage))
		})
	}
}

func TestGetTaskComment(t *testing.T) {
	tests := []struct {
		name      string
		commentID string
		wantErr   bool
	}{
		{
			name:      "get existing comment",
			commentID: "712b8a41-2351-4286-ad03-086eaee4c417",
			wantErr:   false,
		},
		{
			name:      "get non-existent comment",
			commentID: "non-existent-id",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := logic.GetTaskComment(context.Background(), tt.commentID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, comment)
			assert.Equal(t, tt.commentID, comment.Id)
		})
	}
}
