package fitdump

import (
	"encoding/json"
	"fmt"
	"time"
)

// WeightLog is slice of type []WeightEntry, it defines what you should expect
// to find in a typical Fitbit Data Export file for weight data.
type WeightLog []WeightEntry

// RawWeightEntry is an exported weight data collection point from Fitbit in the
// exact format they use, e.g. separate date and time fields. Since the Go
// standard library does not have a "civil date" and "civil time" type, these
// are left as raw strings so you can do as you wish with them.
type RawWeightEntry struct {
	ID     int     `json:"logId"`
	Weight float64 `json:"weight"`
	BMI    float64 `json:"bmi"`
	Fat    float64 `json:"fat"`
	Source string  `json:"source"`
	Date   string  `json:"date"`
	Time   string  `json:"time"`
}

// RecordedAt returns a timestamp in the local timezone corresponding to the
// underlying Date and Time string fields.
func (rw *RawWeightEntry) RecordedAt() (time.Time, error) {
	return time.ParseInLocation(
		"01/02/06 15:04:05",
		fmt.Sprintf("%s %s", rw.Date, rw.Time),
		time.Local,
	)
}

// WeightEntry is a single data collection point for a body weight recording.
type WeightEntry struct {
	ID         int       `json:"logId"`
	Weight     float64   `json:"weight"`
	BMI        float64   `json:"bmi"`
	Fat        float64   `json:"fat"`
	Source     string    `json:"source"`
	RecordedAt time.Time `json:"recorded_at"`
}

// UnmarshalJSON implements json.Unmarshaler
func (we *WeightEntry) UnmarshalJSON(data []byte) error {
	var rwe RawWeightEntry
	err := json.Unmarshal(data, &rwe)
	if err != nil {
		return fmt.Errorf("failed to unmarshal underlying data: %w", err)
	}

	ts, err := rwe.RecordedAt()
	if err != nil {
		return fmt.Errorf("failed to decode underlying timestamp: %w", err)
	}

	we.ID = rwe.ID
	we.Weight = rwe.Weight
	we.BMI = rwe.BMI
	we.Fat = rwe.Fat
	we.Source = rwe.Source
	we.RecordedAt = ts
	return nil
}
