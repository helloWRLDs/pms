package models

type SprintSummary struct {
	Title       string
	Description string
	StartDate   string
	EndDate     string
	TotalTasks  int
	DoneTasks   int
	UndoneTasks int
}

func (s SprintSummary) TemplateName() string {
	return "report"
}
