package render

var (
	_ Renderable = GreetContent{}
	_ Renderable = TaskAssignmentContent{}
)

type GreetContent struct {
	Name        string
	CompanyName string
}

type TaskAssignmentContent struct {
	AssigneeName string
	TaskName     string
	TaskId       string
	ProjectName  string
	CompanyName  string
}

func (c GreetContent) Template() string {
	return "greet.html"
}

func (c GreetContent) Subject() string {
	return "Welcome to " + c.CompanyName
}

func (c TaskAssignmentContent) Template() string {
	return "task_assignment.html"
}

func (c TaskAssignmentContent) Subject() string {
	return "New Task Assignment: " + c.TaskName
}
