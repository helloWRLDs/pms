package logic

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	projectclient "pms.analytics/internal/clients/project"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

type mockListProjectClient struct {
	pb.ProjectServiceClient
	sprint *dto.Sprint
	tasks  []*dto.Task
}

func (m *mockListProjectClient) GetSprint(ctx context.Context, req *pb.GetSprintRequest, opts ...grpc.CallOption) (*pb.GetSprintResponse, error) {
	return &pb.GetSprintResponse{
		Sprint: m.sprint,
	}, nil
}

func (m *mockListProjectClient) ListTasks(ctx context.Context, req *pb.ListTasksRequest, opts ...grpc.CallOption) (*pb.ListTasksResponse, error) {
	return &pb.ListTasksResponse{
		Tasks: &dto.TaskList{
			Items:      m.tasks,
			TotalItems: int32(len(m.tasks)),
		},
	}, nil
}

func TestGetSprint(t *testing.T) {
	now := time.Now()
	sprintID := "test-sprint"
	projectID := "test-project"

	sprint := &dto.Sprint{
		Id:          sprintID,
		ProjectId:   projectID,
		Title:       "Test Sprint",
		Description: "Test Sprint Description",
		StartDate:   timestamppb.New(now),
		EndDate:     timestamppb.New(now.Add(14 * 24 * time.Hour)),
	}

	tasks := []*dto.Task{
		{
			Id:        "task1",
			Type:      string(consts.TaskTypeFeature),
			Priority:  1,
			Status:    string(consts.TASK_STATUS_DONE),
			CreatedAt: timestamppb.New(now.Add(-24 * time.Hour)),
			UpdatedAt: timestamppb.New(now),
			DueDate:   timestamppb.New(now.Add(24 * time.Hour)),
		},
		{
			Id:        "task2",
			Type:      string(consts.TaskTypeBug),
			Priority:  2,
			Status:    string(consts.TASK_STATUS_IN_PROGRESS),
			CreatedAt: timestamppb.New(now.Add(-12 * time.Hour)),
			UpdatedAt: timestamppb.New(now),
		},
	}

	mockListProjectClient := &mockListProjectClient{
		sprint: sprint,
		tasks:  tasks,
	}

	l := &Logic{
		log:           zap.NewNop().Sugar(),
		projectClient: &projectclient.ProjectClient{ProjectServiceClient: mockListProjectClient},
	}

	got, err := l.getSprint(context.Background(), sprintID)
	if err != nil {
		t.Fatalf("getSprint() error = %v", err)
	}

	if got.Id != sprintID {
		t.Errorf("getSprint().Id = %v, want %v", got.Id, sprintID)
	}

	if got.ProjectId != projectID {
		t.Errorf("getSprint().ProjectId = %v, want %v", got.ProjectId, projectID)
	}

	if len(got.Tasks) != len(tasks) {
		t.Errorf("getSprint().Tasks length = %v, want %v", len(got.Tasks), len(tasks))
	}

	for i, task := range got.Tasks {
		if task.Id != tasks[i].Id {
			t.Errorf("getSprint().Tasks[%d].Id = %v, want %v", i, task.Id, tasks[i].Id)
		}
	}
}
