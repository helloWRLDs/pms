package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Greeter(t *testing.T) {
	err := Notifier.GreetUser("Bob", "danil.li24x@gmail.com")
	assert.NoError(t, err)
}
