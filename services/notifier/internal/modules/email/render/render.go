package render

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"pms.pkg/errs"
)

var (
	templates = map[string]string{
		"render.GreetContent": greetTemplate,
	}
	subjects = map[string]string{
		"render.GreetContent": "Greeting",
	}
)

func Render(value any) ([]byte, error) {
	rt := reflect.TypeOf(value).String()

	templateContent, ok := templates[rt]
	if !ok {
		return nil, errs.ErrInternal{
			Reason: "failed finding email template",
		}
	}

	t, err := template.New("email").Parse(templateContent)
	if err != nil {
		return nil, errs.ErrInternal{
			Reason: "failed parsing email template",
		}
	}

	var buf bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	subject, ok := subjects[rt]
	if !ok {
		subject = "No Subject"
	}
	buf.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, headers)))

	if err := t.Execute(&buf, value); err != nil {
		return nil, errs.ErrInternal{
			Reason: "failed executing email template",
		}
	}

	return buf.Bytes(), nil
}
