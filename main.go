package main

import (
	app "app/src"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(app.Width, app.Height)
	ebiten.SetWindowTitle("Змейка на Go")
	ebiten.SetTPS(15)
	ebiten.RunGame(app.New())
}
