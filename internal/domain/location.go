package domain

import (
	"fmt"
	"strings"
	"subscription-bot/internal/geocoding"

	"github.com/sirupsen/logrus"
)

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewLocation(lat, lon float64) Location {
	return Location{Latitude: lat, Longitude: lon}
}

func (l Location) Decode() (string, error) {
	loc, err := geocoding.NewGeocodingAPI().DecodeCoordinates(l.Latitude, l.Longitude)
	if err != nil {
		logrus.Warnf("failed to decode location: %v", err)
	}
	return loc, err
}

func (l Location) String() string {
	decoded, err := l.Decode()
	if err != nil {
		return strings.ReplaceAll(fmt.Sprintf("%v,%v", l.Latitude, l.Longitude), ".", "\\.")
	}
	return strings.ReplaceAll(decoded, ".", "\\.")
}
