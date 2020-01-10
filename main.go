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
	circle       = 'O'
	cross        = 'X'
	h            = windowHeight / 3
	w            = windowWidth / 3
)

func drawBoard() *imdraw.IMDraw {
	boardImDraw := imdraw.New(nil)
	boardImDraw.Color = colornames.Black

	lines := [4][2]pixel.Vec{
		{pixel.V(w, 0), pixel.V(w, windowHeight)},
		{pixel.V(w*2, 0), pixel.V(w*2, windowHeight)},
		{pixel.V(0, h), pixel.V(windowWidth, h)},
		{pixel.V(0, h*2), pixel.V(windowWidth, h*2)},
	}
	for _, linePoints := range lines {
		boardImDraw.Push(linePoints[0], linePoints[1])
		boardImDraw.Line(3)
	}

	return boardImDraw
}

func drawBoardState(state [3][3]string) *imdraw.IMDraw {
	stateImDraw := imdraw.New(nil)
	stateImDraw.Color = colornames.Black

	for lineIndex, line := range state {
		for columnIndex, spot := range line {
			var x float64 = w*float64(lineIndex) + w/2
			var y float64 = h*float64(columnIndex) + h/2

			if spot == string(circle) {
				stateImDraw.Push(
					pixel.V(x, y),
				)
				stateImDraw.Circle(45, 3)
			} else {
				var xr float64 = w / 4
				stateImDraw.Push(
					pixel.V(x-xr, y-xr),
					pixel.V(x+xr, y+xr),
				)
				stateImDraw.Line(3)
				stateImDraw.Push(
					pixel.V(x+xr, y-xr),
					pixel.V(x-xr, y+xr),
				)
				stateImDraw.Line(3)
			}
		}
	}

	return stateImDraw
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

	state := [3][3]string{
		{"X", "O", "O"},
		{"X", "O", "X"},
		{"X", "O", "X"},
	}
	boardImDraw := drawBoard()
	stateImDraw := drawBoardState(state)

	for !win.Closed() {
		win.Clear(colornames.White)
		boardImDraw.Draw(win)
		stateImDraw.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
