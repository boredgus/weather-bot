package geocoding

import (
	"encoding/json"
	"fmt"
	"subscription-bot/config"

	geo "github.com/kellydunn/golang-geo"
	"github.com/sirupsen/logrus"
)

type Geocode struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type GeocodingAPI struct {
	geocoder *geo.GoogleGeocoder
}

func NewGeocodingAPI() GeocodingAPI {
	return GeocodingAPI{
		geocoder: new(geo.GoogleGeocoder),
	}
}

func (api GeocodingAPI) DecodeCoordinates(lat, lon float64) (string, error) {
	bytes, err := api.geocoder.Request(fmt.Sprintf("latlng=%f,%f&key=%v", lat, lon, config.GetConfig().GoogleMapsAPIKey))
	if err != nil {
		logrus.Warnf("failed to fetch data from geocoder: %v", err)
		return "", err
	}

	var res Geocode
	if err := json.Unmarshal(bytes, &res); err != nil {
		logrus.Warnf("failed to unmarshal response: %v", err)
		return "", err
	}

	if res.Status != "OK" {
		return "", fmt.Errorf(res.ErrorMessage)
	}
	if len(res.Results) > 2 {
		return res.Results[2].FormattedAddress, nil
	}
	return res.Results[len(res.Results)-1].FormattedAddress, nil
}
