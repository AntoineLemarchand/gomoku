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
)

type Game struct {
	vector *vector.Path
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		board.placeStone(x, y)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	board.draw(screen)
	x, y := ebiten.CursorPosition()
	board.previewStone(x, y, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return (boardSize*cellSize + boardPad*2), (boardSize*cellSize + boardPad*2)
}

func main() {
	ebiten.SetWindowSize((boardSize+boardPad*2)*cellSize, (boardSize+boardPad*2)*cellSize)
	ebiten.SetWindowTitle("Gomoku")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
