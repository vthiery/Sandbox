package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	WIDTH  = 800
	HEIGHT = 600
	N      = 200
	LENGTH = 25
)

type lines struct {
	abs   [N]float64
	ord   [N]float64
	speed [N]float64
}

func randomAbscissa() float64 {
	return rand.Float64() * WIDTH
}

func randomSpeed() float64 {
	return 10.0 * (1.0 + rand.Float64())
}

func newLines() lines {
	var l lines
	for i := 0; i < N; i++ {
		l.abs[i] = randomAbscissa()
		l.ord[i] = rand.Float64() * HEIGHT // generate ordinates randomly to mimick warm-up
		l.speed[i] = randomSpeed()
	}
	return l
}

func (l *lines) update() {
	for i := 0; i < N; i++ {
		if l.ord[i] < 0.0 {
			l.abs[i] = randomAbscissa()
			l.ord[i] = HEIGHT
			l.speed[i] = randomSpeed()
		}
	}
}

func (l *lines) draw(win *pixelgl.Window) {
	for i := 0; i < N; i++ {
		l.ord[i] = l.ord[i] - l.speed[i]
		imd := imdraw.New(nil)
		imd.Color = colornames.Gainsboro
		imd.EndShape = imdraw.RoundEndShape
		imd.Push(pixel.V(l.abs[i], l.ord[i]), pixel.V(l.abs[i], l.ord[i]-LENGTH))
		imd.Line(3)
		imd.Draw(win)
	}
}

func gameLoop() {
	// setup the main window
	cfg := pixelgl.WindowConfig{
		Title:  "Rain!",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// start the random generator
	rand.Seed(123)

	// run the actual game loop
	l := newLines()
	for !win.Closed() {
		win.Clear(colornames.Lightskyblue)

		l.draw(win)
		l.update()

		win.Update()
	}
}

func main() {
	pixelgl.Run(gameLoop)
}
