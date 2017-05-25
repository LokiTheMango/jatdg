package game

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/input"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
)

// Game Object
type Game struct {
	screen        Screen
	level         level.Level
	input         input.Keyboard
	DrawRequested bool
	camera        entities.Mob
}

//Constructor
func New() *Game {
	game := &Game{
		screen:        Screen{},
		input:         input.Keyboard{},
		DrawRequested: false,
	}
	return game
}

func (game *Game) Init(filePath string) {
	game.screen = NewScreen(filePath)
	game.level = level.NewLevel(game.screen.SpriteSheet, enums.LEVEL_WIDTH, enums.LEVEL_HEIGHT)
	game.screen.SetLevel(game.level)
	game.camera = entities.NewCamera()
}

func (game *Game) Update() {
	game.camera.Update()
	game.moveObjects()
	game.clearScreen()
	game.render()
}

func (game *Game) render() {
	x := game.camera.GetX()
	y := game.camera.GetY()
	game.screen.RenderLevel(x, y)
}

func (game *Game) clearScreen() {
	game.screen.ClearScreen()
}

func (game *Game) moveObjects() {
	xOffset := 0
	yOffset := 0
	if game.input.Up {
		yOffset--
	}
	if game.input.Down {
		yOffset++
	}
	if game.input.Left {
		xOffset--
	}
	if game.input.Right {
		xOffset++
	}
	if xOffset != 0 || yOffset != 0 {
		game.camera.Move(xOffset, yOffset)
	}
}

func (game *Game) UpdateInput(newInput input.Keyboard) {
	game.input = newInput
}

func (game *Game) GetSpriteSheet() render.SpriteSheet {
	return game.screen.SpriteSheet
}

func (game *Game) GetPixelArray() []byte {
	return game.screen.PixelArray
}
