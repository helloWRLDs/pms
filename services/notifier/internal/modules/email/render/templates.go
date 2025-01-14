package render

import (
	"embed"
	"fmt"

	"pms.pkg/errs"
)

var (
	//go:embed docs/*.html
	htmlTemplates embed.FS
)

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
