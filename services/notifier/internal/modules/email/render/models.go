package render

type Renderable interface {
	Template() string
	Subject() string
}

type GreetContent struct {
	Renderable
	CompanyName string
	Name        string
}

func (gc GreetContent) Template() string {
	return "greet.html"
}

func (gc GreetContent) Subject() string {
	return "Greeting"
}
