package logic

import (
	"context"
	"os"
	"testing"
)

func Test_DownloadFile(t *testing.T) {
	docID := "326b474d-caaf-44c6-8d2c-149996e0b7ac"
	doc, err := logic.DownloadDocument(context.Background(), docID)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("./test.pdf", doc.Body, 0644)
}
