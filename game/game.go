package game

// Game Object
type Game struct {
	gameMap       Map
	input         Input
	DrawRequested bool
}

//Constructor
func New() *Game {
	return &Game{
		gameMap:       Map{},
		input:         Input{},
		DrawRequested: false,
	}
}

func (game *Game) Update() {

}

func (game *Game) UpdateInput(newInput Input) {
	game.input = newInput
}
