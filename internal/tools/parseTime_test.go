package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var valueToResultTime = map[string]struct {
	input string
	time  time.Time
	err   error
}{
	"invalid_time_format_1": {
		input: "tyu",
		time:  time.Now(),
		err:   InvaTimeFormat,
	},
	"invalid_time_format_2": {
		input: "1:1",
		time:  time.Now(),
		err:   InvaTimeFormat,
	},
	"latitude_is_not_a_number": {
		input: "tyu,we",
		time:  time.Now(),
		err:   InvaTimeFormat,
	},
	"valid_time_format_1": {
		input: "01:01",
		time:  time.UnixMilli(-62167215540000),
		err:   nil,
	},
	"valid_time_format_2": {
		input: "20:20",
		time:  time.UnixMilli(-62167146000000),
		err:   nil,
	},
}

func TestMessageToTime(t *testing.T) {
	for name, value := range valueToResultTime {
		time, err := MessageToTime(value.input)

		assert.Equal(t, value.time.Unix(), time.Unix(), name)
		if value.err != nil {
			assert.ErrorContains(t, err, value.err.Error(), name)
			continue
		}
		assert.Equal(t, value.err, err, name)
	}
}
