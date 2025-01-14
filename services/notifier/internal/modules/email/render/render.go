package render

import (
	"bytes"
	"fmt"
	"text/template"

	"pms.pkg/errs"
)

func Render(value Renderable) ([]byte, error) {
	tmpl, err := getTemplate(value.Template())
	if err != nil {
		return []byte{}, err
	}

	t, err := template.New("email").Parse(string(tmpl))
	if err != nil {
		return nil, errs.ErrInternal{
			Reason: "failed parsing email template",
		}
	}

	var buf bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	buf.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", value.Subject(), headers)))

	if err := t.Execute(&buf, value); err != nil {
		return nil, errs.ErrInternal{
			Reason: "failed executing email template",
		}
	}

	return buf.Bytes(), nil
}
