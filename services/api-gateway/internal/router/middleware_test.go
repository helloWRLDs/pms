package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
)

// MockAuthClient is a mock implementation of the auth client
type MockAuthClient struct {
	mock.Mock
}

func (m *MockAuthClient) GetUserRole(ctx context.Context, userID, companyID string) (*dto.Role, error) {
	args := m.Called(ctx, userID, companyID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.Role), args.Error(1)
}

func TestRequirePermission(t *testing.T) {
	app := fiber.New()
	mockAuth := new(MockAuthClient)

	// Test cases
	tests := []struct {
		name           string
		role           *dto.Role
		permission     string
		expectedStatus int
	}{
		{
			name: "CEO has all permissions",
			role: &dto.Role{
				Name: "ceo",
				Permissions: []string{
					string(consts.COMPANY_READ_PERMISSION),
					string(consts.COMPANY_WRITE_PERMISSION),
					string(consts.USER_READ_PERMISSION),
					string(consts.USER_WRITE_PERMISSION),
					string(consts.PROJECT_READ_PERMISSION),
					string(consts.PROJECT_WRITE_PERMISSION),
					string(consts.TASK_READ_PERMISSION),
					string(consts.TASK_WRITE_PERMISSION),
					string(consts.SPRINT_READ_PERMISSION),
					string(consts.SPRINT_WRITE_PERMISSION),
				},
			},
			permission:     string(consts.COMPANY_WRITE_PERMISSION),
			expectedStatus: http.StatusOK,
		},
		{
			name: "Project Manager has limited permissions",
			role: &dto.Role{
				Name: "project_manager",
				Permissions: []string{
					string(consts.PROJECT_READ_PERMISSION),
					string(consts.PROJECT_WRITE_PERMISSION),
					string(consts.TASK_READ_PERMISSION),
					string(consts.TASK_WRITE_PERMISSION),
					string(consts.SPRINT_READ_PERMISSION),
					string(consts.SPRINT_WRITE_PERMISSION),
				},
			},
			permission:     string(consts.COMPANY_WRITE_PERMISSION),
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Participant has minimal permissions",
			role: &dto.Role{
				Name: "participant",
				Permissions: []string{
					string(consts.PROJECT_READ_PERMISSION),
					string(consts.TASK_READ_PERMISSION),
					string(consts.TASK_WRITE_PERMISSION),
					string(consts.SPRINT_READ_PERMISSION),
				},
			},
			permission:     string(consts.PROJECT_WRITE_PERMISSION),
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test route
			app.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("success")
			})

			// Setup mock expectations
			mockAuth.On("GetUserRole", mock.Anything, "test-user", "test-company").Return(tt.role, nil)

			// Create test request
			req := httptest.NewRequest("GET", "/test", nil)
			req.Header.Set("X-User-ID", "test-user")
			req.Header.Set("X-Company-ID", "test-company")

			// Create test response
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestRequirePermissionErrorHandling(t *testing.T) {
	app := fiber.New()
	mockAuth := new(MockAuthClient)

	tests := []struct {
		name           string
		mockError      error
		expectedStatus int
	}{
		{
			name:           "Database error",
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Role not found",
			mockError:      nil,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test route
			app.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("success")
			})

			// Setup mock expectations
			mockAuth.On("GetUserRole", mock.Anything, "test-user", "test-company").Return(nil, tt.mockError)

			// Create test request
			req := httptest.NewRequest("GET", "/test", nil)
			req.Header.Set("X-User-ID", "test-user")
			req.Header.Set("X-Company-ID", "test-company")

			// Create test response
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestRequirePermissionMissingHeaders(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		name           string
		headers        map[string]string
		expectedStatus int
	}{
		{
			name:           "Missing User ID",
			headers:        map[string]string{"X-Company-ID": "test-company"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing Company ID",
			headers:        map[string]string{"X-User-ID": "test-user"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing both headers",
			headers:        map[string]string{},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test route
			app.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("success")
			})

			// Create test request
			req := httptest.NewRequest("GET", "/test", nil)
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			// Create test response
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
