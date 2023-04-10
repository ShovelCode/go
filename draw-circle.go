package main

import (
    "math"

    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/vg"
    "github.com/gonum/plot/vg/draw"
)

func main() {
    // Create a new plot and set its title and labels
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Circle"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    // Create a plotter for the circle
    n := 100 // Number of points to plot
    pts := make(plotter.XYs, n)
    for i := range pts {
        theta := float64(i) * 2 * math.Pi / float64(n)
        pts[i].X = math.Cos(theta)
        pts[i].Y = math.Sin(theta)
    }
    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    line.Color = draw.DarkBlue

    // Add the plotter to the plot
    p.Add(line)

    // Save the plot to a file
    if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "circle.png"); err != nil {
        panic(err)
    }
}
