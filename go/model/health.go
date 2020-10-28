package model

import (
	"fmt"
	"strings"
	"time"
)

type (
	Health struct {
		BloodSugar int      `json:"blood_sugar"`
		Ts         JSONTime `json:"ts"`
	}

	JSONTime time.Time
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = JSONTime(time.Time{})
		return nil
	}
	defaultTime, err := time.Parse(timeLayout, s)
	*t = JSONTime(defaultTime)

	return err
}
