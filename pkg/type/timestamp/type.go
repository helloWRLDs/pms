package timestamp

import "time"

type Timestamp struct {
	time.Time
	Format TimeFormat
}

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp{
		Time:   t,
		Format: DEFAULT_FORMAT,
	}
}

func WithFormat(t time.Time, format TimeFormat) Timestamp {
	return Timestamp{
		Time:   t,
		Format: format,
	}
}

func (ts Timestamp) String() string {
	if ts.Format == "" {
		ts.Format = DEFAULT_FORMAT
	}
	return ts.Time.Format(string(ts.Format))
}
