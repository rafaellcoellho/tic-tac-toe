package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowHeight = 400
	windowWidth  = 400
	circle       = "O"
	cross        = "X"
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

func drawBoardState(state [3][3]string, crossColor, circleColor color.RGBA) *imdraw.IMDraw {
	stateImDraw := imdraw.New(nil)

	for lineIndex, line := range state {
		for columnIndex, spot := range line {
			// make 0,0 top left, not bottom left
			var y float64 = windowWidth - (w*float64(lineIndex) + w/2)
			var x float64 = h*float64(columnIndex) + h/2

			if spot == circle {
				stateImDraw.Color = circleColor
				stateImDraw.Push(
					pixel.V(x, y),
				)
				stateImDraw.Circle(45, 3)
			} else if spot == cross {
				stateImDraw.Color = crossColor
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

func blankStateExist(state [3][3]string) bool {
	for _, line := range state {
		for _, spot := range line {
			if spot == "" {
				return true
			}
		}
	}
	return false
}

func checkWinner(state [3][3]string) string {
	winner := ""

	for i := 0; i < 3; i++ {
		if state[i][0] == state[i][1] && state[i][1] == state[i][2] {
			winner = state[i][0]
			break
		}
	}

	for i := 0; i < 3; i++ {
		if state[0][i] == state[1][i] && state[1][i] == state[2][i] {
			winner = state[0][i]
			break
		}
	}

	if state[0][0] == state[1][1] && state[1][1] == state[2][2] {
		winner = state[0][0]
	}

	if state[0][2] == state[1][1] && state[1][1] == state[2][0] {
		winner = state[0][2]
	}

	if winner == "" && !blankStateExist(state) {
		return "tie"
	}

	return winner
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
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	currentPlayer := cross

	boardImDraw := drawBoard()
	result := ""
	circleColor := colornames.Black
	crossColor := colornames.Black

	for !win.Closed() {
		if result == "" && win.JustPressed(pixelgl.MouseButtonLeft) {
			x := win.MousePosition().X
			y := win.MousePosition().Y

			var columnIndex int
			var lineIndex int

			if x < w {
				columnIndex = 0
			} else if x > w && x < w*2 {
				columnIndex = 1
			} else {
				columnIndex = 2
			}

			if y < h {
				lineIndex = 2
			} else if y > h && y < h*2 {
				lineIndex = 1
			} else {
				lineIndex = 0
			}

			if state[lineIndex][columnIndex] == "" {
				state[lineIndex][columnIndex] = currentPlayer
				if currentPlayer == cross {
					currentPlayer = circle
				} else {
					currentPlayer = cross
				}
			}
		}

		win.Clear(colornames.White)
		boardImDraw.Draw(win)
		stateImDraw := drawBoardState(state, crossColor, circleColor)
		stateImDraw.Draw(win)
		win.Update()

		result = checkWinner(state)
		if result != "" {
			if result == "X" {
				crossColor = colornames.Red
			} else if result == "O" {
				circleColor = colornames.Red
			} else {
				crossColor = colornames.Green
				circleColor = colornames.Green
			}
		}
	}
}

func main() {
	pixelgl.Run(run)
}
