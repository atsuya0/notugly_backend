package domain

import (
	"encoding/json"
	"time"
)

const (
	layout = "2006-01-02 15:04:05"
)

type JsonTime struct {
	Data time.Time
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + j.Data.Format(layout) + "\""), nil
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	var strTime string
	if err := json.Unmarshal(b, &strTime); err != nil {
		return err
	}
	time, err := time.Parse(layout, strTime)
	if err != nil {
		return err
	}

	j.Data = time
	return nil
}
