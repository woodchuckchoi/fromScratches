package model

import (
	"fmt"
	"time"
)

type (
	Health struct {
		BloodSugar int      `json:"blood_sugar"`
		Ts         JSONTime `json:"ts"`
	}

	JSONTime time.Time
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%s", time.Time(t).Format("2020-10-26 20:22:38"))
	return []byte(stamp), nil
}
