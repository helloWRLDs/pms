package models

type SprintSummary struct {
	Title       string
	Description string
	StartDate   string
	EndDate     string
	TotalTasks  int
	DoneTasks   int
	UndoneTasks int

	// Enhanced metrics
	TotalPoints     int32
	CompletionRate  float64
	TasksByType     map[string]int
	TasksByPriority map[string]int
	UserPerformance []UserPerformance
	TopPerformers   []UserPerformance
	TaskInsights    TaskInsights
	SprintVelocity  int32
}

type UserPerformance struct {
	UserID         string
	FirstName      string
	LastName       string
	FullName       string
	DoneTasks      int32
	TotalTasks     int32
	TotalPoints    int32
	CompletionRate float64
}

type TaskInsights struct {
	AveragePointsPerTask float64
	HighestValueTask     string
	MostCommonType       string
	MostCommonPriority   string
	TasksCompletedOnTime int
	TasksOverdue         int
}

func (s SprintSummary) TemplateName() string {
	return "report"
}
