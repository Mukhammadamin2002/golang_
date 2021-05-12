package main

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	graph := chart.BarChart{
		Title: "World Pollution",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height:   256,
		BarWidth: 50,
		Bars: []chart.Value{
			{Value: 65, Label: "1950"},
			{Value: 50, Label: "1975"},
			{Value: 75, Label: "1990"},
			{Value: 69, Label: "2010"},
			{Value: 80, Label: "2020"},
		},
	}
	f, _ := os.Create("barchart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
