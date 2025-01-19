package timestamp

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

func (ts Timestamp) Value() (driver.Value, error) {
	return ts.String(), nil
}

func (ts *Timestamp) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case time.Time:
		ts.Time = v
		ts.Format = DEFAULT_FORMAT
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	for _, format := range formats {
		if parsed, err := time.Parse(string(format), str); err == nil {
			ts.Time = parsed
			ts.Format = format
			return nil
		}
	}
	return fmt.Errorf("failed to parse timestamp: %s", str)
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(ts.String())
}

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to unmarshal timestamp: %w", err)
	}

	for _, format := range formats {
		if parsed, err := time.Parse(string(format), str); err == nil {
			ts.Time = parsed
			ts.Format = format
			return nil
		}
	}

	return fmt.Errorf("failed to parse timestamp: %s", str)
}
