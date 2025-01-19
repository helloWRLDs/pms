package render

type Renderable interface {
	Template() string
	Subject() string
}

type RenderInfo struct {
	template string
	subject  string
}
