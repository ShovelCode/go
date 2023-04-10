package main

import (
	"math"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

func main() {
	// Define the function that represents the lemniscate
	f := func(t float64) float64 {
		return math.Sqrt(2) * math.Cos(t) / (math.Sin(t)*math.Sin(t) + 1)
	}

	// Create a new plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Set the title and labels
	p.Title.Text = "Lemniscate"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// Create a plotter for the lemniscate
	pts := make(plotter.XYs, 500)
	for i := range pts {
		t := float64(i) / 100.0
		pts[i].X = f(t)
		pts[i].Y = f(t) * math.Sin(t)
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	line.Color = draw.DarkBlue

	// Add the plotter to the plot
	p.Add(line)

	// Save the plot to a file
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "lemniscate.png"); err != nil {
		panic(err)
	}
}
