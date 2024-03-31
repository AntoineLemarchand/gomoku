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
