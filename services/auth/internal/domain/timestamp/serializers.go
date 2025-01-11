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
	switch v := value.(type) {
	case time.Time:
		*ts = Timestamp(v)
	case string:
		parsed, err := time.Parse(SQLITE_FORMAT, v)
		if err != nil {
			return fmt.Errorf("failed to parse timestamp: %w", err)
		}
		*ts = Timestamp(parsed)
	case []byte:
		parsed, err := time.Parse(SQLITE_FORMAT, string(v))
		if err != nil {
			return fmt.Errorf("failed to parse timestamp: %w", err)
		}
		*ts = Timestamp(parsed)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	t := time.Time(ts)
	return json.Marshal(t.Format(SQLITE_FORMAT))
}

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to unmarshal timestamp: %w", err)
	}

	parsed, err := time.Parse(SQLITE_FORMAT, str)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}
	*ts = Timestamp(parsed)
	return nil
}
