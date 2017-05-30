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
	Width  int
	Height int

	spriteSheet render.SpriteSheet
	levelSheet  render.LevelSheet
	Tiles       []tiles.Tile
}

func NewLevel(spriteSheet render.SpriteSheet, filePath string) Level {
	rand.Seed(time.Now().UTC().UnixNano())
	levelSheet, width, height := render.NewLevelSheet(filePath)
	level := Level{
		spriteSheet: spriteSheet,
		levelSheet:  levelSheet,
		Width:       width,
		Height:      height,
	}
	level.generateLevel()
	return level
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (level *Level) generateLevel() {
	level.Tiles = make([]tiles.Tile, level.Height*level.Width)
	for i := 0; i < level.Height; i++ {
		for j := 0; j < level.Width; j++ {
			//// random Tile generation for Tests
			//nextTile := randInt(0, 2)
			nextTile := 1
			indexHeight := i * level.Width * 4
			pix := level.levelSheet.PixelArray[indexHeight+j*4 : indexHeight+(j+1)*4]
			wall := []byte{255, 255, 255, 255}
			if testEq(pix, wall) {
				nextTile = 0
			}
			tileType := enums.TileType(nextTile)
			level.Tiles[i*level.Width+j] = tiles.NewTile(j, i, tileType, level.spriteSheet.PixelArray)
		}
	}
}

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}