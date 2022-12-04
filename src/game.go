package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Game struct {
	snake *Snake
	food  *Food
	dir   Dir
	input KeyboardInput
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	screen.Fill(color.RGBA{A: 255, R: 21, G: 21, B: 20})
	for i := 0; i < Cols; i++ {
		for j := 0; j < Rows; j++ {
			DrawCell(screen, i, j, color.Black, Gap, CellSize)
		}
	}
}

func New() *Game {
	return &Game{}
}
