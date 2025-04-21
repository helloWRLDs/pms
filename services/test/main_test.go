package main

import (
	"os"
	"testing"
)

func Test_ReadPNG(t *testing.T) {
	f, err := os.ReadFile("./user_profile.png")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(f))
	os.WriteFile("./user_profile_raw.txt", f, 0644)
}
