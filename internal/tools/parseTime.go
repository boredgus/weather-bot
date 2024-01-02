package tools

import (
	"errors"
	"fmt"
	"time"
)

const TimeTemplate = "15:04"

var InvaTimeFormat = errors.New("failed to parse time")

func MessageToTime(msg string) (t time.Time, e error) {
	t, e = time.Parse(TimeTemplate, msg)
	if e != nil {
		return time.Now(), fmt.Errorf("%v: %w", InvaTimeFormat, e)
	}
	return t, nil
}
