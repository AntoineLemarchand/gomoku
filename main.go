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

type Game struct {
	vector *vector.Path
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var board = newBoard(boardSize, boardPad, cellSize, 2)
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
