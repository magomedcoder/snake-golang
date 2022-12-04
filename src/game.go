package app

type Game struct {
	snake *Snake
	food  *Food
	dir   Dir
	input KeyboardInput
}

func New() *Game {
	return &Game{}
}
