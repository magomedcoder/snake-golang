package app

import "image/color"

type Food struct {
	Cell
}

func (f *Food) Update(x, y int) {
	f.x = x
	f.y = y
}

func NewFood() *Food {
	return &Food{
		Cell: Cell{
			x: Cols / 2,
			y: Rows / 2,
			c: color.RGBA{R: 85, G: 165, B: 34, A: 255},
		},
	}
}
