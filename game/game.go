package game

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/game/render"
)

// Game Object
type Game struct {
	level           Level
	input           Input
	DrawRequested   bool
	SpriteSheet     render.SpriteSheet
	SpriteSheetSize int
	PixelArray      []byte
	xOffset         int
	yOffset         int
}

//Constructor
func New() *Game {
	rand.Seed(time.Now().UTC().UnixNano())
	game := &Game{
		level:         Level{},
		input:         Input{},
		DrawRequested: false,
		SpriteSheet:   render.SpriteSheet{},
		xOffset:       0,
		yOffset:       0,
	}
	return game
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (game *Game) Init(filePath string) {
	game.createSpriteSheet(filePath)
	game.createLevel()
	game.parseFrameBuffer()
}

func (game *Game) createLevel() {
	game.level = NewLevel(game.SpriteSheet, 10, 10)
}

func (game *Game) parseFrameBuffer() {
	game.PixelArray = game.level.ParseFrameBuffer(game.SpriteSheetSize, game.xOffset, game.yOffset)
}

func (game *Game) Update() {
	game.checkInputForOffsets()
	game.parseFrameBuffer()
}

func (game *Game) checkInputForOffsets() {
	if game.input.Up {
		game.yOffset++
	}
	if game.input.Down {
		game.yOffset--
	}
	if game.input.Left {
		game.xOffset++
	}
	if game.input.Right {
		game.xOffset--
	}
}

func (game *Game) UpdateInput(newInput Input) {
	game.input = newInput
}

func (game *Game) createSpriteSheet(filePath string) {
	spriteSheet, size := render.NewSpriteSheet(filePath)
	game.SpriteSheet = spriteSheet
	game.SpriteSheetSize = size
}
func (game *Game) GetSpriteMap() render.SpriteSheet {
	return game.SpriteSheet
}

func (game *Game) GetPixelArray() []byte {
	return game.SpriteSheet.PixelArray
}
