package weather

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"subscription-bot/internal/tools"
	"time"

	"github.com/wcharczuk/go-chart"
)

type WeatherForecastItem struct {
	Title []struct {
		Details string `json:"description"`
	} `json:"weather"`
	Temperature struct {
		Actual float64 `json:"temp"`
	} `json:"main"`
	Date          int64   `json:"dt"`
	Precipitation float64 `json:"pop"`
}

func (i WeatherForecastItem) description() string {
	var b strings.Builder
	for _, w := range i.Title {
		fmt.Fprintf(&b, "%v, ", w.Details)
	}
	return b.String()
}

func (i WeatherForecastItem) String() string {
	return fmt.Sprintf(
		"%v <i>%v</i><b>%v℃</b> &#9730; %v%%\n",
		time.UnixMilli(i.Date*1000).Format("15:04"),
		i.description(),
		math.Round(i.Temperature.Actual),
		math.Round(i.Precipitation*100),
	)
}

type WeatherForecast struct {
	List []WeatherForecastItem `json:"list"`
}

func NewWeatherForecast(listItem ...WeatherForecastItem) WeatherForecast {
	return WeatherForecast{List: listItem}
}

func (f WeatherForecast) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for _, i := range f.List {
		buffer.Write([]byte(i.String()))
	}
	return buffer.String()
}

func (g WeatherForecast) getForecastChartValues() (
	xValues []float64,
	yValues []float64,
	xTicks []chart.Tick,
	yTicks []chart.Tick,
) {
	xValues = make([]float64, len(g.List))
	yValues = make([]float64, len(g.List))
	xTicks = make([]chart.Tick, len(g.List))

	for idx, f := range g.List {
		xValues[idx] = float64(idx)
		yValues[idx] = f.Temperature.Actual
		xTicks[idx] = chart.Tick{
			Value: float64(idx),
			Label: time.UnixMilli(f.Date * 1000).Format("15:04"),
		}
	}

	minT, maxT, err := tools.FindMinAndMax(yValues)
	if err != nil {
		return
	}
	yTicks = make([]chart.Tick, maxT-minT+4)
	for idx := range yTicks {
		v := minT - 1 + idx
		yTicks[idx] = chart.Tick{
			Value: float64(v),
			Label: fmt.Sprintf("%v℃", v),
		}
	}
	return
}

func (g WeatherForecast) generateChart() chart.Chart {
	xValues, yValues, xTicks, yTicks := g.getForecastChartValues()
	return chart.Chart{
		Width: 500,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlack,
			},
			TickStyle: chart.StyleShow(),
			Ticks:     xTicks,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlack,
			},
			TickStyle: chart.StyleShow(),
			Ticks:     yTicks,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}
}

func (f WeatherForecast) RenderGraph() ([]byte, error) {
	var buffer bytes.Buffer
	err := f.generateChart().Render(chart.PNG, &buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to render graph: %w", err)
	}
	return buffer.Bytes(), nil
}
