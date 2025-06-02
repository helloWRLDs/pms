package logic

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	authclient "pms.analytics/internal/clients/auth"
	projectclient "pms.analytics/internal/clients/project"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

type mockAuthClient struct {
	pb.AuthServiceClient
	users []*dto.User
}

func (m *mockAuthClient) ListUsers(ctx context.Context, req *pb.ListUsersRequest, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
	return &pb.ListUsersResponse{
		UserList: &dto.UserList{
			Items: m.users,
		},
	}, nil
}

type mockProjectClient struct {
	pb.ProjectServiceClient
	projects []*dto.Project
	sprints  map[string][]*dto.Sprint
	tasks    map[string][]*dto.Task
}

func (m *mockProjectClient) ListProjects(ctx context.Context, req *pb.ListProjectsRequest, opts ...grpc.CallOption) (*pb.ListProjectsResponse, error) {
	return &pb.ListProjectsResponse{
		Projects: &dto.ProjectList{
			Items: m.projects,
		},
	}, nil
}

func (m *mockProjectClient) ListSprints(ctx context.Context, req *pb.ListSprintsRequest, opts ...grpc.CallOption) (*pb.ListSprintsResponse, error) {
	sprints := m.sprints[req.Filter.ProjectId]
	return &pb.ListSprintsResponse{
		Sprints: &dto.SprintList{
			Items: sprints,
		},
	}, nil
}

func (m *mockProjectClient) ListTasks(ctx context.Context, req *pb.ListTasksRequest, opts ...grpc.CallOption) (*pb.ListTasksResponse, error) {
	tasks := m.tasks[req.Filter.SprintId]
	if tasks == nil {
		tasks = m.tasks["overall"]
	}
	return &pb.ListTasksResponse{
		Tasks: &dto.TaskList{
			Items:      tasks,
			TotalItems: int32(len(tasks)),
		},
	}, nil
}

func TestGetUserTaskStats(t *testing.T) {
	now := time.Now()
	companyID := "test-company"
	projectID := "test-project"
	sprintID := "test-sprint"
	userID := "test-user"

	users := []*dto.User{
		{
			Id:        userID,
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	projects := []*dto.Project{
		{
			Id: projectID,
		},
	}

	sprints := map[string][]*dto.Sprint{
		projectID: {
			{
				Id: sprintID,
			},
		},
	}

	tasks := map[string][]*dto.Task{
		sprintID: {
			{
				Id:         "task1",
				Type:       string(consts.TaskTypeFeature),
				Priority:   1,
				Status:     string(consts.TASK_STATUS_DONE),
				AssigneeId: userID,
				CreatedAt:  timestamppb.New(now.Add(-24 * time.Hour)),
				UpdatedAt:  timestamppb.New(now),
				DueDate:    timestamppb.New(now.Add(24 * time.Hour)),
			},
			{
				Id:         "task2",
				Type:       string(consts.TaskTypeBug),
				Priority:   2,
				Status:     string(consts.TASK_STATUS_IN_PROGRESS),
				AssigneeId: userID,
				CreatedAt:  timestamppb.New(now.Add(-12 * time.Hour)),
				UpdatedAt:  timestamppb.New(now),
			},
			{
				Id:         "task3",
				Type:       string(consts.TaskTypeStory),
				Priority:   3,
				Status:     string(consts.TASK_STATUS_CREATED),
				AssigneeId: userID,
				CreatedAt:  timestamppb.New(now),
			},
		},
		"overall": {
			{
				Id:         "task4",
				Type:       string(consts.TaskTypeChore),
				Priority:   4,
				Status:     string(consts.TASK_STATUS_DONE),
				AssigneeId: userID,
				CreatedAt:  timestamppb.New(now.Add(-48 * time.Hour)),
				UpdatedAt:  timestamppb.New(now),
			},
		},
	}

	mockAuthClient := &mockAuthClient{users: users}
	mockProjectClient := &mockProjectClient{
		projects: projects,
		sprints:  sprints,
		tasks:    tasks,
	}

	l := &Logic{
		log:           zap.NewNop().Sugar(),
		authClient:    &authclient.AuthClient{AuthServiceClient: mockAuthClient},
		projectClient: &projectclient.ProjectClient{ProjectServiceClient: mockProjectClient},
	}

	stats, err := l.GetUserTaskStats(context.Background(), companyID)
	if err != nil {
		t.Fatalf("GetUserTaskStats() error = %v", err)
	}

	if len(stats) != 1 {
		t.Fatalf("expected 1 user stats, got %d", len(stats))
	}

	userStats := stats[0]
	if userStats.UserId != userID {
		t.Errorf("user ID = %v, want %v", userStats.UserId, userID)
	}

	sprintStats := userStats.Stats[sprintID]
	if sprintStats == nil {
		t.Fatal("sprint stats not found")
	}

	if sprintStats.TotalTasks != 3 {
		t.Errorf("sprint total tasks = %v, want 3", sprintStats.TotalTasks)
	}
	if sprintStats.DoneTasks != 1 {
		t.Errorf("sprint done tasks = %v, want 1", sprintStats.DoneTasks)
	}
	if sprintStats.InProgressTasks != 1 {
		t.Errorf("sprint in progress tasks = %v, want 1", sprintStats.InProgressTasks)
	}
	if sprintStats.ToDoTasks != 1 {
		t.Errorf("sprint todo tasks = %v, want 1", sprintStats.ToDoTasks)
	}

	overallStats := userStats.Stats["overall"]
	if overallStats == nil {
		t.Fatal("overall stats not found")
	}

	if overallStats.TotalTasks != 1 {
		t.Errorf("overall total tasks = %v, want 1", overallStats.TotalTasks)
	}
	if overallStats.DoneTasks != 1 {
		t.Errorf("overall done tasks = %v, want 1", overallStats.DoneTasks)
	}
}

func TestGetPriorityName(t *testing.T) {
	tests := []struct {
		name     string
		priority int32
		want     string
	}{
		{
			name:     "lowest priority",
			priority: 1,
			want:     "Lowest",
		},
		{
			name:     "low priority",
			priority: 2,
			want:     "Low",
		},
		{
			name:     "medium priority",
			priority: 3,
			want:     "Medium",
		},
		{
			name:     "high priority",
			priority: 4,
			want:     "High",
		},
		{
			name:     "highest priority",
			priority: 5,
			want:     "Highest",
		},
		{
			name:     "unknown priority",
			priority: 6,
			want:     "Unknown",
		},
		{
			name:     "zero priority",
			priority: 0,
			want:     "Unknown",
		},
		{
			name:     "negative priority",
			priority: -1,
			want:     "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{}
			if got := l.getPriorityName(tt.priority); got != tt.want {
				t.Errorf("getPriorityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMostCommon(t *testing.T) {
	tests := []struct {
		name   string
		counts map[string]int
		want   string
	}{
		{
			name: "single item",
			counts: map[string]int{
				"a": 1,
			},
			want: "a",
		},
		{
			name: "multiple items with clear winner",
			counts: map[string]int{
				"a": 3,
				"b": 1,
				"c": 2,
			},
			want: "a",
		},
		{
			name: "tie for most common",
			counts: map[string]int{
				"a": 2,
				"b": 2,
				"c": 1,
			},
			want: "a",
		},
		{
			name:   "empty map",
			counts: map[string]int{},
			want:   "",
		},
		{
			name: "all equal counts",
			counts: map[string]int{
				"a": 1,
				"b": 1,
				"c": 1,
			},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{}
			if got := l.findMostCommon(tt.counts); got != tt.want {
				t.Errorf("findMostCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
