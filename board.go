package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type board struct {
	size      int
	pad       int
	cellSize  int
	lineWidth int
}

func newBoard(size, pad, cellSize, lineWidth int) *board {
	return &board{size, pad, cellSize, lineWidth}
}

func (b *board) draw(screen *ebiten.Image) {
	screen.Fill(colornames.Amber100)

	for i := 0; i <= b.size; i++ {
		x := float32(i*b.cellSize) + float32(b.pad*b.cellSize)
		y := float32(b.pad * b.cellSize)
		vector.StrokeLine(screen, x, y, x, y+float32(b.size*b.cellSize), float32(b.lineWidth), colornames.Grey900, true)
		vector.StrokeLine(screen, y, x, y+float32(b.size*b.cellSize), x, float32(b.lineWidth), colornames.Grey900, true)
	}
}

func (b *board) cellAt(x, y int) (int, int) {
	return (x - b.pad*b.cellSize) / b.cellSize, (y - b.pad*b.cellSize) / b.cellSize
}

func (b *board) cellCenter(x, y int) (int, int) {
	return x*b.cellSize + b.cellSize/2 + b.pad*b.cellSize, y*b.cellSize + b.cellSize/2 + b.pad*b.cellSize
}

func (b *board) previewStone(x, y int, screen *ebiten.Image) {
	cell_x_id, cell_y_id := b.cellAt(x, y)
	if (cell_x_id >= 0 && cell_x_id < boardSize) && (cell_y_id >= 0 && cell_y_id < boardSize) {
		cell_x, cell_y := b.cellCenter(cell_x_id, cell_y_id)
		vector.DrawFilledCircle(screen, float32(cell_x), float32(cell_y), cellSize/2-2, colornames.Green400, true)
	}
}
