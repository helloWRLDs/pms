package models

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Models(t *testing.T) {
	ts := timestamppb.New(time.Now())
	t.Logf("%#v", ts)
}
