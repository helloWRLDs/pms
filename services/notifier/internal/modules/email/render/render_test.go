package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
	greet := NewGreetContent("Bob", "AITU")
	t.Log("subject: ", greet.Subject())
	t.Log("template: ", greet.Template())

	tmpl, err := getTemplate(greet.Template())

	assert.NotNil(t, tmpl)
	assert.NoError(t, err)
	t.Log("template length: ", len(tmpl))
}
