package companydata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddParticipant(t *testing.T) {
	err := repo.AddParticipant(context.Background(), "bca9bbdf-b9e5-4bfd-b024-49c06dc56227", "92a77fc1-23eb-4d84-b649-d810dd21174f")
	assert.NoError(t, err)
}
