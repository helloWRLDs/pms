package render

import (
	_ "embed"
)

var (
	//go:embed docs/greet.html
	greetTemplate string
)
