package render

var (
	_ Renderable = greetContent{}
)

type greetContent struct {
	CompanyName string
	Name        string

	RenderInfo
}

func NewGreetContent(name, companyName string) Renderable {
	return &greetContent{
		Name:        name,
		CompanyName: companyName,
		RenderInfo: RenderInfo{
			subject:  "Greeting",
			template: "greet.html",
		},
	}
}

func (gc greetContent) Template() string {
	return gc.template
}

func (gc greetContent) Subject() string {
	return gc.subject
}
