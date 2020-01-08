package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowHeight = 400
	windowWidth  = 400
)

func drawBoard() *imdraw.IMDraw {
	var w float64 = windowWidth / 3
	var h float64 = windowHeight / 3

	board := imdraw.New(nil)
	board.Color = colornames.Black

	lines := [4][2]pixel.Vec{
		{pixel.V(w, 0), pixel.V(w, windowHeight)},
		{pixel.V(w*2, 0), pixel.V(w*2, windowHeight)},
		{pixel.V(0, h), pixel.V(windowWidth, h)},
		{pixel.V(0, h*2), pixel.V(windowWidth, h*2)},
	}
	for _, linePoints := range lines {
		board.Push(linePoints[0], linePoints[1])
		board.Line(3)
	}
	return board
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "tic-tac-toe",
		Bounds: pixel.R(0, 0, windowHeight, windowWidth),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	board := drawBoard()

	for !win.Closed() {
		win.Clear(colornames.White)

		board.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
