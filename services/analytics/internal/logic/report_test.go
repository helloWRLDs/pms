package logic

import (
	"context"
	"testing"

	htmlmodule "pms.analytics/internal/modules/htmlgen"
	"pms.analytics/internal/modules/htmlgen/models"
	"pms.pkg/transport/grpc/dto"
)

func TestGenerateSprintReport(t *testing.T) {
	projectID := "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1"
	sprintID := "f46ecc84-09cb-47b8-b8f5-b9c1198bc4d9"

	report, err := logic.CreateReportTemplate(context.Background(), &dto.DocumentCreation{
		Title:     "Sprint Report",
		ProjectId: projectID,
		SprintId:  sprintID,
	})
	if err != nil {
		t.Fatalf("failed to generate sprint report: %v", err)
	}
	t.Log(report)
}

func TestSprintReportTemplate(t *testing.T) {
	sprintSummary := models.SprintSummary{
		Title:           "Test Sprint",
		Description:     "Test sprint description",
		StartDate:       "January 1, 2024",
		EndDate:         "January 15, 2024",
		TotalTasks:      10,
		DoneTasks:       8,
		UndoneTasks:     2,
		TotalPoints:     150,
		CompletionRate:  80.0,
		SprintVelocity:  150,
		TasksByType:     map[string]int{"Feature": 5, "Bug": 3, "Story": 2},
		TasksByPriority: map[string]int{"High": 4, "Medium": 4, "Low": 2},
		UserPerformance: []models.UserPerformance{
			{
				UserID:         "user1",
				FirstName:      "John",
				LastName:       "Doe",
				FullName:       "John Doe",
				DoneTasks:      5,
				TotalTasks:     6,
				TotalPoints:    80,
				CompletionRate: 83.3,
			},
			{
				UserID:         "user2",
				FirstName:      "Jane",
				LastName:       "Smith",
				FullName:       "Jane Smith",
				DoneTasks:      3,
				TotalTasks:     4,
				TotalPoints:    70,
				CompletionRate: 75.0,
			},
		},
		TopPerformers: []models.UserPerformance{
			{
				UserID:         "user1",
				FirstName:      "John",
				LastName:       "Doe",
				FullName:       "John Doe",
				DoneTasks:      5,
				TotalTasks:     6,
				TotalPoints:    80,
				CompletionRate: 83.3,
			},
		},
		TaskInsights: models.TaskInsights{
			AveragePointsPerTask: 15.0,
			HighestValueTask:     "Implement authentication",
			MostCommonType:       "Feature",
			MostCommonPriority:   "High",
		},
	}

	htmlData := htmlmodule.Template[models.SprintSummary]{
		Title:   "Test Sprint Report",
		Name:    models.SprintSummary{}.TemplateName(),
		Content: sprintSummary,
	}

	html, err := htmlmodule.Render(htmlData)
	if err != nil {
		t.Fatalf("failed to render template: %v", err)
	}

	if len(html) == 0 {
		t.Fatal("rendered HTML is empty")
	}

	t.Logf("Template rendered successfully, HTML length: %d bytes", len(html))
}

func TestSprintReportTemplateComparisons(t *testing.T) {
	testCases := []struct {
		name           string
		completionRate float64
		undoneTasks    int
		expectHigh     bool
		expectLow      bool
	}{
		{
			name:           "High completion rate",
			completionRate: 85.5,
			undoneTasks:    1,
			expectHigh:     true,
			expectLow:      false,
		},
		{
			name:           "Low completion rate",
			completionRate: 65.2,
			undoneTasks:    5,
			expectHigh:     false,
			expectLow:      true,
		},
		{
			name:           "Exactly 80% completion",
			completionRate: 80.0,
			undoneTasks:    2,
			expectHigh:     true,
			expectLow:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sprintSummary := models.SprintSummary{
				Title:          "Test Sprint",
				Description:    "Test sprint description",
				StartDate:      "January 1, 2024",
				EndDate:        "January 15, 2024",
				TotalTasks:     10,
				DoneTasks:      8,
				UndoneTasks:    tc.undoneTasks,
				TotalPoints:    150,
				CompletionRate: tc.completionRate,
				TaskInsights: models.TaskInsights{
					AveragePointsPerTask: 15.0,
					HighestValueTask:     "Test task",
					MostCommonType:       "Feature",
					MostCommonPriority:   "High",
				},
			}

			htmlData := htmlmodule.Template[models.SprintSummary]{
				Title:   "Test Sprint Report",
				Name:    models.SprintSummary{}.TemplateName(),
				Content: sprintSummary,
			}

			html, err := htmlmodule.Render(htmlData)
			if err != nil {
				t.Fatalf("failed to render template: %v", err)
			}

			htmlStr := string(html)

			if tc.expectHigh {
				if !contains(htmlStr, "High completion rate") {
					t.Errorf("Expected high completion rate message for %.1f%%", tc.completionRate)
				}
			}

			if tc.expectLow {
				if !contains(htmlStr, "could be higher") {
					t.Errorf("Expected low completion rate message for %.1f%%", tc.completionRate)
				}
			}

			t.Logf("Test case '%s' passed with completion rate %.1f%%", tc.name, tc.completionRate)
		})
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		(len(s) >= len(substr)) &&
		func() bool {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}()
}

func TestPrintSampleReport(t *testing.T) {
	t.Skip("Run manually to see sample report output")

	sprintSummary := models.SprintSummary{
		Title:           "Q4 Sprint 3",
		Description:     "Focus on user authentication and payment integration",
		StartDate:       "December 1, 2024",
		EndDate:         "December 15, 2024",
		TotalTasks:      15,
		DoneTasks:       12,
		UndoneTasks:     3,
		TotalPoints:     245,
		CompletionRate:  80.0,
		SprintVelocity:  245,
		TasksByType:     map[string]int{"Feature": 8, "Bug": 4, "Story": 2, "Chore": 1},
		TasksByPriority: map[string]int{"High": 5, "Medium": 7, "Low": 3},
		UserPerformance: []models.UserPerformance{
			{
				UserID:         "user1",
				FirstName:      "Alice",
				LastName:       "Johnson",
				FullName:       "Alice Johnson",
				DoneTasks:      8,
				TotalTasks:     9,
				TotalPoints:    135,
				CompletionRate: 88.9,
			},
			{
				UserID:         "user2",
				FirstName:      "Bob",
				LastName:       "Smith",
				FullName:       "Bob Smith",
				DoneTasks:      4,
				TotalTasks:     6,
				TotalPoints:    110,
				CompletionRate: 66.7,
			},
		},
		TopPerformers: []models.UserPerformance{
			{
				UserID:         "user1",
				FirstName:      "Alice",
				LastName:       "Johnson",
				FullName:       "Alice Johnson",
				DoneTasks:      8,
				TotalTasks:     9,
				TotalPoints:    135,
				CompletionRate: 88.9,
			},
		},
		TaskInsights: models.TaskInsights{
			AveragePointsPerTask: 16.3,
			HighestValueTask:     "Implement OAuth2 authentication",
			MostCommonType:       "Feature",
			MostCommonPriority:   "Medium",
		},
	}

	htmlData := htmlmodule.Template[models.SprintSummary]{
		Title:   "Q4 Sprint 3 Report",
		Name:    models.SprintSummary{}.TemplateName(),
		Content: sprintSummary,
	}

	html, err := htmlmodule.Render(htmlData)
	if err != nil {
		t.Fatalf("failed to render template: %v", err)
	}

	t.Logf("Sample Sprint Report HTML:\n%s", string(html))
}
