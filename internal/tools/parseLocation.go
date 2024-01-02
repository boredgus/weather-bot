package tools

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"subscription-bot/internal/domain"
)

var InvalidLocationFormat = errors.New("failed to parse time from message")
var InvalidLatitudeFormat = errors.New("failed to parse latitude value")
var InvalidLatitudeRange = errors.New("provided value for latitude does not satisfies [-90,90]")
var InvalidLongitudeFormat = errors.New("failed to parse longitude value")
var InvalidLongitudeRange = errors.New("provided value for longitude does not satisfies [-180,180]")

func MessageToLocation(msg string) (domain.Location, error) {
	coords := strings.Split(msg, ",")
	if len(coords) != 2 {
		return domain.Location{}, InvalidLocationFormat
	}

	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return domain.Location{}, InvalidLatitudeFormat
	}
	if math.Abs(lat) > 90 {
		return domain.Location{}, InvalidLatitudeRange
	}

	lon, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return domain.Location{}, InvalidLongitudeFormat
	}
	if math.Abs(lon) > 180 {
		return domain.Location{}, InvalidLongitudeRange
	}

	return domain.Location{Latitude: lat, Longitude: lon}, nil
}
