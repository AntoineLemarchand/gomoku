package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

const (
	COLOR_BLACK = 1
	COLOR_WHITE = 2
)

type stone struct {
	x, y  int
	color int
}

type Board struct {
	size      int
	pad       int
	cellSize  int
	lineWidth int
	stones    []stone
	turn      int
}

func newBoard(size, pad, cellSize, lineWidth int) *Board {
	return &Board{size, pad, cellSize, lineWidth, []stone{}, COLOR_BLACK}
}

func (b *Board) draw(screen *ebiten.Image) {
	screen.Fill(colornames.Amber100)
	var stroke_length = (b.size - 1) * b.cellSize

	for i := 0; i < b.size; i++ {
		offset := float32(b.pad) + 1./2*float32(b.cellSize)
		x := float32(i*b.cellSize) + offset
		y := offset
		vector.StrokeLine(screen, x, y, x, y+float32(stroke_length), float32(b.lineWidth), colornames.Grey900, true)
		vector.StrokeLine(screen, y, x, y+float32(stroke_length), x, float32(b.lineWidth), colornames.Grey900, true)
	}

	for _, s := range b.stones {
		x, y := b.cellCenter(s.x, s.y)
		var color = colornames.Black
		if s.color == COLOR_WHITE {
			color = colornames.White
		}
		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(b.cellSize/2-2), color, true)
	}
}

func (b *Board) cellAt(x, y int) (int, int) {
	x -= b.pad
	y -= b.pad
	if x < 0 || y < 0 {
		return -1, -1
	}
	return x / b.cellSize, y / b.cellSize
}

func (b *Board) cellCenter(x, y int) (int, int) {
	return b.pad + x*b.cellSize + b.cellSize/2, b.pad + y*b.cellSize + b.cellSize/2
}

func (b *Board) isCellTaken(x, y int) bool {
	for _, s := range b.stones {
		if s.x == x && s.y == y {
			return true
		}
	}
	return false
}

func (b *Board) placeStone(x, y int) {
	if b.isCellTaken(x, y) {
		return
	}
	x, y = b.cellAt(x, y)
	if (x < 0 || x >= b.size) || (y < 0 || y >= b.size) {
		return
	}
	b.stones = append(b.stones, stone{x, y, b.turn})
	if b.turn == COLOR_BLACK {
		b.turn = COLOR_WHITE
	} else {
		b.turn = COLOR_BLACK
	}
}

func (b *Board) previewStone(x, y int, screen *ebiten.Image) {
	cell_x_id, cell_y_id := b.cellAt(x, y)
	if (cell_x_id >= 0 && cell_x_id < boardSize) && (cell_y_id >= 0 && cell_y_id < boardSize) {
		cell_x, cell_y := b.cellCenter(cell_x_id, cell_y_id)
		var preview_color = colornames.Blue400
		if b.isCellTaken(cell_x_id, cell_y_id) {
			preview_color = colornames.Red400
		}
		vector.DrawFilledCircle(screen, float32(cell_x), float32(cell_y), cellSize/2-2, preview_color, true)
	}
}
