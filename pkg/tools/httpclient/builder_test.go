package httpclient

import (
	"encoding/json"
	"testing"
)

func Test_Build(t *testing.T) {
	user := struct {
		Name string `json:"name"`
	}{
		Name: "Bob",
	}
	res := New().
		Method("GET").
		URL("https://api.kanye.rest").
		Headers("Content-Type", "application/json").
		Query(
			"page", 1,
			"per_page", 2,
			"ads",
		).
		Body(user)
	t.Log(res.String())
}

func Test_Response(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	bob := User{
		Name: "Bob",
	}
	m, err := json.Marshal(bob)
	if err != nil {
		t.Fatal(err)
	}
	res := Response{
		Data:   m,
		Status: 200,
	}
	var received User
	if err := res.ScanJSON(&received); err != nil {
		t.Fatal(err)
	}
	t.Log(received)
}
