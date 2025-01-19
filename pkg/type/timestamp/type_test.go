package timestamp

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Serialize(t *testing.T) {
	ts := WithFormat(time.Now(), ISO_FORMAT)
	j, err := json.Marshal(ts)
	if !assert.NoError(t, err) {
		return
	}
	t.Log(string(j))

	var unmarshalled Timestamp
	json.Unmarshal(j, &unmarshalled)
	if !assert.NoError(t, err) {
		return
	}
	t.Logf("%#v", unmarshalled)
	t.Log(unmarshalled.Format)
}

func Test_Time(t *testing.T) {
	ts := time.Now()
	t.Log(ts)
}
