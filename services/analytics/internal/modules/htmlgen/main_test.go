package htmlmodule

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FS(t *testing.T) {
	fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		assert.NoError(t, err)
		t.Log(path)
		return nil
	})
}

func Test_Render(t *testing.T) {
	type Content struct {
		Text string
	}
	content := Template[Content]{
		Title: "some content",
		Content: Content{
			Text: "some text",
		},
		Name: "report",
	}
	data, err := Render(content)
	assert.NoError(t, err)
	t.Log(string(data))
}

func Test_HTMLtoPDF(t *testing.T) {
	html := `1 <p>some text</p><p><img alt="Mathieu 'ZywOo' Herbaut" src="https://img-cdn.hltv.org/playerbodyshot/Xkqvuwl9o12Mi20Vd0lzHl.png?bg=3e4c54&amp;h=200&amp;ixlib=java-2.1.0&amp;rect=124%2C8%2C467%2C467&amp;w=200&amp;s=7b46e9daaa5aede9c8f77058755f9932" class="picture" title="Mathieu 'ZywOo' Herbaut"><br></p>`
	PDF([]byte(html))
}
