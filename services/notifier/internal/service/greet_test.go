package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Greeter(t *testing.T) {
	err := Notifier.GreetUser(context.Background(), "Danil", "danil.li24x@gmail.com")
	assert.NoError(t, err)
}
