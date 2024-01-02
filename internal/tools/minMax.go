package tools

import (
	"errors"
)

var EmptySliceError = errors.New("given an empty slice")

func FindMinAndMax(numbers []float64) (int, int, error) {
	if len(numbers) < 1 {
		return 0, 0, EmptySliceError
	}

	min, max := numbers[0], numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return int(min), int(max), nil
}
