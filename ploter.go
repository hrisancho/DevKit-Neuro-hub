package main

import (
	"container/ring"
	"image"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"github.com/wcharczuk/go-chart/seq"
)

const w, h = 960, 540

func main() {
	s := ebiten.DeviceScaleFactor()

	g := NewGraph(100, s, func(g *Graph) {
		g.Background = chart.Style{
			Padding: chart.Box{
				Top:    25,
				Left:   25,
				Right:  25,
				Bottom: 25,
			},
			FillColor: drawing.ColorFromHex("efefef"),
		}

		g.YAxis = chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Max: 4,
				Min: -4,
			},
		}
	})

	go generateRandomData(g, 64*time.Millisecond)

	m := g.Image()

	go func() {
		for range time.NewTicker(64 * time.Millisecond).C {
			m = g.Image()
		}
	}()

	update := func(screen *ebiten.Image) error {
		if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
			os.Exit(0)
		}

		if !ebiten.IsDrawingSkipped() {
			screen.ReplacePixels(m.Pix)
		}

		return nil
	}

	if err := ebiten.Run(update, int(w*s), int(h*s), 1/s, "Ebiten + go-chart"); err != nil {
		log.Fatal(err)
	}
}

type Graph struct {
	chart.Chart
	*ring.Ring
}

func NewGraph(n int, s float64, options ...func(*Graph)) *Graph {
	g := &Graph{chart.Chart{
		Width:  int(float64(w) * s),
		Height: int(float64(h) * s),
	},
		ring.New(n),
	}

	for i := 0; i < g.Len(); i++ {
		g.Value = 0.0
		g.Ring = g.Next()
	}

	for _, o := range options {
		o(g)
	}

	return g
}

func (g *Graph) Set(v float64) {
	g.Value = v
	g.Ring = g.Next()
}

func (g *Graph) Image() *image.RGBA {
	var yValues []float64

	g.Do(func(x interface{}) {
		if x != nil {
			yValues = append(yValues, x.(float64))
		}
	})

	mainSeries := chart.ContinuousSeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(0).WithAlpha(96),
			FillColor:   chart.GetDefaultColor(0).WithAlpha(32),
		},
		XValues: seq.Range(1.0, float64(len(yValues))),
		YValues: yValues,
	}

	smaSeries := &chart.SMASeries{
		InnerSeries: mainSeries,
	}

	g.Series = []chart.Series{mainSeries, smaSeries}

	iw := &chart.ImageWriter{}

	g.Render(chart.PNG, iw)

	m, _ := iw.Image()

	return m.(*image.RGBA)
}

func generateRandomData(g *Graph, d time.Duration) {
	for range time.NewTicker(d).C {
		v := g.Value.(float64)

		switch {
		case v > 2:
			g.Set(v - rand.Float64())
		case v < -2:
			g.Set(v + rand.Float64())
		default:
			g.Set(v + 2*rand.Float64() - 0.5)
		}
	}
}
