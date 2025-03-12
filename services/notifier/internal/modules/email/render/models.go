package render

var (
	_ Renderable = GreetContent{}
)

type GreetContent struct {
	CompanyName string
	Name        string
}

func (gc GreetContent) Template() string {
	return "greet.html"
}

func (gc GreetContent) Subject() string {
	return "Greeting"
}
