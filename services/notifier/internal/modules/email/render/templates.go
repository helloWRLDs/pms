package render

import (
	"embed"
	"fmt"

	"pms.pkg/errs"
)

var (
	//go:embed docs/*.html
	htmlTemplates embed.FS

	//go:embed docs/greet.html
	greetTemplate string

	//go:embed docs/task_assignment.html
	taskAssignmentTemplate string
)

var templates = map[string]string{
	"greet.html":           greetTemplate,
	"task_assignment.html": taskAssignmentTemplate,
}

func getTemplate(template string) ([]byte, error) {
	data, err := htmlTemplates.ReadFile(fmt.Sprintf("docs/%s", template))
	if err != nil {
		return []byte{}, errs.ErrNotFound{
			Object: "template",
			Field:  "name",
			Value:  template,
		}
	}
	return data, nil
}
