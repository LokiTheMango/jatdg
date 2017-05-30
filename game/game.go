package game

import (
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/input"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

// Game Object
type Game struct {
	screen        Screen
	level         level.Level
	input         input.Keyboard
	DrawRequested bool
	Towers        []entities.Entity
	Enemies       []entities.Mob
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
	game.screen = NewScreen(filePath + "tiles.png")
	level, tiles := level.NewLevel(game.screen.SpriteSheet, filePath+"level.png")
	game.level = level
	game.screen.SetLevel(&game.level)
	game.camera = entities.NewCamera(&game.level)
	game.createTowerEntities(tiles)
}

func (game *Game) createTowerEntities(tiles []*tiles.Tile) {
	game.Towers = make([]entities.Entity, game.level.NumTowers)
	for i, _ := range game.Towers {
		game.Towers[i] = entities.NewTower(tiles[i])
	}
}

func (game *Game) Update() {
	for _, tower := range game.Towers {
		tower.Update()
	}
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
	//Move enemies here
	//game.moveWithCollisionCheck(xOffset, yOffset)
}

func (game *Game) moveWithCollisionCheck(xa int, ya int) {
	for _, enemy := range game.Enemies {
		x := (enemy.GetX() + xa) / 32
		y := (enemy.GetY()*game.level.Width + ya) / 32
		tileType := game.level.Tiles[x+y].TileType
		if tileType == 0 {
			enemy.Move(xa, ya)
		}
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
