package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	boardSize = 15
	boardPad  = 2
	cellSize  = 40
)

var (
	board = newBoard(boardSize, boardPad, cellSize, 2)
	turn  = COLOR_BLACK
)

type Game struct {
	vector *vector.Path
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= boardPad*cellSize && x < (boardSize+boardPad)*cellSize && y >= boardPad*cellSize && y < (boardSize+boardPad)*cellSize {
			x, y = board.cellAt(x, y)
			board.placeStone(x, y)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	board.draw(screen)
	x, y := ebiten.CursorPosition()
	board.previewStone(x, y, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return (boardSize + boardPad*2) * cellSize, (boardSize + boardPad*2) * cellSize
}

func main() {
	ebiten.SetWindowSize((boardSize+boardPad*2)*cellSize, (boardSize+boardPad*2)*cellSize)
	ebiten.SetWindowTitle("Gomoku")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
