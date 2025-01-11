package timestamp

import (
	"time"
)

const SQLITE_FORMAT = "2006-01-02 15:04:05"

type Timestamp time.Time

func (ts Timestamp) String() string {
	t := time.Time(ts)
	return t.Format(SQLITE_FORMAT)
}
