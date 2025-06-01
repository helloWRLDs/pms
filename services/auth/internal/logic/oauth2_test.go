package logic

import "testing"

func TestInitiateOAuth2(t *testing.T) {
	url, err := logic.InitiateOAuth2("google")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}
