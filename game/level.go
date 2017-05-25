package game

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/render"
)

//Level Object
type Level struct {
	width  int
	height int

	spriteSheet render.SpriteSheet
	tileSize    int
	tiles       []entities.Tile
}

func NewLevel(spriteSheet render.SpriteSheet, width int, height int) Level {
	level := Level{
		spriteSheet: spriteSheet,
		width:       width,
		height:      height,
	}
	level.generateLevel()
	return level
}

func (level *Level) generateLevel() {
	level.tiles = make([]entities.Tile, level.height*level.width)
	for i := 0; i < level.height; i++ {
		for j := 0; j < level.width; j++ {
			//// random Tile generation for Tests
			//nextTile := randInt(0, 4)
			level.tiles[i+j] = entities.NewTile(i, j, enums.VOID, level.spriteSheet.PixelArray)
		}
	}
}

func (level *Level) ParseFrameBuffer(size int, xOffset int, yOffset int) []byte {
	framebuffer := make([]byte, size)
	for y := 0; y < enums.HEIGHT; y++ {
		yOff := y + yOffset
		if yOff < 0 || yOff >= enums.HEIGHT {
			continue
		}
		// PIXEL WIDTH TIME 4 (RGBA)
		for x := 0; x < enums.WIDTH*4; x++ {
			xOff := x + (xOffset << 2)
			if xOff < 0 || xOff >= enums.WIDTH*4 {
				continue
			}
			tileIndex := ((yOff >> 5) & (level.height - 1)) + ((xOff >> 7) & (level.width - 1))
			tileArr := level.tiles[tileIndex].GetPixelArray()
			tileWidth := enums.WIDTH_TILE
			tileHeight := enums.HEIGHT_TILE
			maskTileWidth := tileWidth - 1
			maskTileHeight := tileHeight - 1
			index := ((xOff & maskTileWidth) % tileWidth) + ((yOff&maskTileHeight)%tileHeight)*tileWidth
			framebuffer[x+y*enums.WIDTH*4] = tileArr[index]
		}
	}
	return framebuffer
}
