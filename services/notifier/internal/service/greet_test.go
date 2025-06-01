package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Greeter(t *testing.T) {
	err := Notifier.GreetUser(context.Background(), "Viktor", "kossinovviktor@gmail.com")
	assert.NoError(t, err)
}
