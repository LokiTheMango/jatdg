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

	SpriteSheet render.SpriteSheet
	Tiles       []entities.Tile
}

func NewLevel(spriteSheet render.SpriteSheet, width int, height int) Level {
	level := Level{
		SpriteSheet: spriteSheet,
		width:       width,
		height:      height,
	}
	level.GenerateLevel()
	return level
}

func (level *Level) GenerateLevel() {
	level.Tiles = make([]entities.Tile, level.height*level.width)
	for i := 0; i < level.height; i++ {
		for j := 0; j < level.width; j++ {
			nextSprite := randInt(0, 4)
			level.Tiles[i+j] = entities.NewTile(level.SpriteSheet.PixelArray, enums.TileType(nextSprite), i, j)
		}
	}
}

func (level *Level) ParseFrameBuffer(size int, xOffset int, yOffset int) []byte {
	framebuffer := make([]byte, size)
	for y := 0; y < enums.HEIGHT; y++ {
		yy := y + yOffset
		if yy < 0 || yy >= enums.HEIGHT {
			continue
		}
		for x := 0; x < enums.WIDTH*4; x++ {
			xx := x + (xOffset << 2)
			if xx < 0 || xx >= enums.WIDTH*4 {
				continue
			}
			tileIndex := ((yy >> 5) & 9) + ((xx >> 7) & 9)
			tileArr := level.Tiles[tileIndex].GetPixelArray()
			index := ((xx & 127) % 128) + ((yy&31)%32)*128
			if index < 0 {
				index *= -1
			}
			framebuffer[x+y*enums.WIDTH*4] = tileArr[index]
		}
	}
	return framebuffer
}
