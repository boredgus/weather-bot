package tools

import (
	"fmt"
	"subscription-bot/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var valueToResultLocation = map[string]struct {
	input    string
	location domain.Location
	err      error
}{
	"invalid_input_string": {
		input:    "tyu",
		location: domain.Location{},
		err:      InvalidLocationFormat,
	},
	"invalid_input_multiple_numbers": {
		input:    "1,1,1",
		location: domain.Location{},
		err:      InvalidLocationFormat,
	},
	"latitude_is_not_a_number": {
		input:    "tyu,we",
		location: domain.Location{},
		err:      InvalidLatitudeFormat,
	},
	"latitude_is_out_of_range_positive": {
		input:    "100,we",
		location: domain.Location{},
		err:      InvalidLatitudeRange,
	},
	"latitude_is_out_of_range_negative": {
		input:    "-100,we",
		location: domain.Location{},
		err:      InvalidLatitudeRange,
	},
	"longitude_is_not_a_number": {
		input:    "10,we",
		location: domain.Location{},
		err:      InvalidLongitudeFormat,
	},
	"longitude_is_out_of_range_positive": {
		input:    "0,200",
		location: domain.Location{},
		err:      InvalidLongitudeRange,
	},
	"longitude_is_out_of_range_negative": {
		input:    "0,-200",
		location: domain.Location{},
		err:      InvalidLongitudeRange,
	},
	"valid_location_positive": {
		input:    "10,10",
		location: domain.Location{Latitude: 10, Longitude: 10},
		err:      nil,
	},
	"valid_location_negative": {
		input:    "-20,-20",
		location: domain.Location{Latitude: -20, Longitude: -20},
		err:      nil,
	},
	"valid_location_mixed": {
		input:    "-20,30",
		location: domain.Location{Latitude: -20, Longitude: 30},
		err:      nil,
	},
}

func TestMessageToLocation(t *testing.T) {
	for name, value := range valueToResultLocation {
		loc, err := MessageToLocation(value.input)
		fmt.Println(value.input)

		assert.Equal(t, value.location, loc, name)
		assert.Equal(t, value.err, err, name)
	}
}
