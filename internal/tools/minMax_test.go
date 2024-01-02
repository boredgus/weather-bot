package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var valuesToResult = map[string]struct {
	input []float64
	min   int
	max   int
	err   error
}{
	"differrent_int_values": {
		input: []float64{1, 7, 3, 9, 0},
		min:   0,
		max:   9,
		err:   nil,
	},
	"equal_int_values": {
		input: []float64{1, 1, 1, 1, 1},
		min:   1,
		max:   1,
		err:   nil,
	},
	"one_int_value": {
		input: []float64{111},
		min:   111,
		max:   111,
		err:   nil,
	},
	"no_values": {
		input: []float64{},
		min:   0,
		max:   0,
		err:   EmptySliceError,
	},
	"different_float_values": {
		input: []float64{1.2, 2.3, -10.9, 20.5},
		min:   -10,
		max:   20,
		err:   nil,
	},
	"equal_float_values": {
		input: []float64{1.2, 1.2, 1.2, 1.2},
		min:   1,
		max:   1,
		err:   nil,
	},
	"one_float_value": {
		input: []float64{2.1},
		min:   2,
		max:   2,
		err:   nil,
	},
}

func TestFindMinAndMax(t *testing.T) {
	for name, value := range valuesToResult {
		min, max, err := FindMinAndMax(value.input)

		assert.Equal(t, value.min, min, name)
		assert.Equal(t, value.max, max, name)
		assert.Equal(t, value.err, err, name)
	}
}
