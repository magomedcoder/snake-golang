package main

import (
	app "app/src"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(app.Width, app.Height)
	ebiten.SetWindowTitle("Змейка на Go")
	ebiten.SetTPS(15)
	if err := ebiten.RunGame(app.New()); err != nil {
		log.Fatal(err)
	}
}
