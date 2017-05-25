package level

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

//Level Object
type Level struct {
	width  int
	height int

	spriteSheet render.SpriteSheet
	Tiles       []tiles.Tile
}

func NewLevel(spriteSheet render.SpriteSheet, width int, height int) Level {
	rand.Seed(time.Now().UTC().UnixNano())
	level := Level{
		spriteSheet: spriteSheet,
		width:       width,
		height:      height,
	}
	level.generateLevel()
	return level
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (level *Level) generateLevel() {
	level.Tiles = make([]tiles.Tile, level.height*level.width)
	for i := 0; i < level.height; i++ {
		for j := 0; j < level.width; j++ {
			//// random Tile generation for Tests
			nextTile := randInt(0, 2)
			tileType := enums.TileType(nextTile)
			level.Tiles[i*level.width+j] = tiles.NewTile(j, i, tileType, level.spriteSheet.PixelArray)
		}
	}
}
