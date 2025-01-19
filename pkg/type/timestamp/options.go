package timestamp

import "time"

type TimeFormat string

const (
	SQLITE_FORMAT  TimeFormat = "2006-01-02 15:04:05"
	ISO_FORMAT     TimeFormat = "2006-01-02T15:04:05.999Z07:00"
	DEFAULT_FORMAT TimeFormat = time.RFC3339
	DATE_ONLY      TimeFormat = "2006-01-02"
	TIME_ONLY      TimeFormat = "15:04:05"
)

func (t TimeFormat) String() string {
	switch t {
	case SQLITE_FORMAT:
		return "SQLITE"
	case ISO_FORMAT:
		return "ISO_FORMAT"
	case DATE_ONLY:
		return "DATE_ONLY"
	case TIME_ONLY:
		return "TIME_ONLY"
	default:
		return "DEFAULT_FORMAT"
	}
}

var (
	formats = []TimeFormat{
		ISO_FORMAT,
		SQLITE_FORMAT,
		DEFAULT_FORMAT,
		DATE_ONLY,
		TIME_ONLY,
	}
)
