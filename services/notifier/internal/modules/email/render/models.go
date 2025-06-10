package render

var (
	_ Renderable = GreetContent{}
	_ Renderable = TaskAssignmentContent{}
	_ Renderable = WelcomeLoginContent{}
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

type LoginDetails struct {
	Timestamp string
	Location  string
	Device    string
	IPAddress string
}

type SocialLinks struct {
	Website  string
	Twitter  string
	LinkedIn string
}

type WelcomeLoginContent struct {
	Name              string
	CompanyName       string
	CompanyAddress    string
	IsFirstLogin      bool
	DashboardURL      string
	SetupGuideURL     string
	SupportEmail      string
	SupportPhone      string
	DocumentationURL  string
	VideoTutorialsURL string
	LoginDetails      *LoginDetails
	SocialLinks       SocialLinks
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

func (c WelcomeLoginContent) Template() string {
	return "welcome_login.html"
}

func (c WelcomeLoginContent) Subject() string {
	if c.IsFirstLogin {
		return "Welcome to " + c.CompanyName + " - Let's Get Started!"
	}
	return "Welcome back to " + c.CompanyName
}
