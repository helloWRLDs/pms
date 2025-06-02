package timestamp

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewTimestamp(t *testing.T) {
	now := time.Now()
	ts := NewTimestamp(now)

	if !ts.Time.Equal(now) {
		t.Errorf("NewTimestamp() = %v, want %v", ts.Time, now)
	}
	if ts.Format != DEFAULT_FORMAT {
		t.Errorf("NewTimestamp() format = %v, want %v", ts.Format, DEFAULT_FORMAT)
	}
}

func TestWithFormat(t *testing.T) {
	now := time.Now()
	ts := WithFormat(now, ISO_FORMAT)

	if !ts.Time.Equal(now) {
		t.Errorf("WithFormat() = %v, want %v", ts.Time, now)
	}
	if ts.Format != ISO_FORMAT {
		t.Errorf("WithFormat() format = %v, want %v", ts.Format, ISO_FORMAT)
	}
}

func TestTimestamp_String(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		ts       Timestamp
		expected string
	}{
		{
			name:     "default format",
			ts:       NewTimestamp(now),
			expected: now.Format(string(DEFAULT_FORMAT)),
		},
		{
			name:     "ISO format",
			ts:       WithFormat(now, ISO_FORMAT),
			expected: now.Format(string(ISO_FORMAT)),
		},
		{
			name:     "date only format",
			ts:       WithFormat(now, DATE_ONLY),
			expected: now.Format(string(DATE_ONLY)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.String(); got != tt.expected {
				t.Errorf("Timestamp.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimestamp_Value(t *testing.T) {
	now := time.Now()
	ts := NewTimestamp(now)

	val, err := ts.Value()
	if err != nil {
		t.Errorf("Timestamp.Value() error = %v", err)
	}
	if val != ts.String() {
		t.Errorf("Timestamp.Value() = %v, want %v", val, ts.String())
	}
}

func TestTimestamp_Scan(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{
			name:    "time.Time value",
			value:   now,
			wantErr: false,
		},
		{
			name:    "string value",
			value:   now.Format(string(ISO_FORMAT)),
			wantErr: false,
		},
		{
			name:    "invalid type",
			value:   42,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ts Timestamp
			err := ts.Scan(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Timestamp.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	now := time.Now()
	ts := NewTimestamp(now)

	data, err := ts.MarshalJSON()
	if err != nil {
		t.Errorf("Timestamp.MarshalJSON() error = %v", err)
	}

	var result string
	if err := json.Unmarshal(data, &result); err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
	}

	if result != ts.String() {
		t.Errorf("Timestamp.MarshalJSON() = %v, want %v", result, ts.String())
	}
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	now := time.Now()
	timeStr := now.Format(string(ISO_FORMAT))
	data, _ := json.Marshal(timeStr)

	var ts Timestamp
	err := ts.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("Timestamp.UnmarshalJSON() error = %v", err)
	}

	if !ts.Time.Equal(now) {
		t.Errorf("Timestamp.UnmarshalJSON() = %v, want %v", ts.Time, now)
	}
}
