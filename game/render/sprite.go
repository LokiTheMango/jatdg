package render

import (
	"github.com/LokiTheMango/jatdg/enums"
)

type Sprite struct {
	tileType   enums.TileType
	PixelArray []byte
}

func NewSprite(pixelArray []byte, tileType enums.TileType, posX int, posY int) Sprite {
	pixel := make([]byte, enums.WIDTH_TILE*enums.HEIGHT_TILE)
	for i := 0; i < enums.HEIGHT_TILE; i++ {
		start := i * enums.WIDTH_TILE
		end := start + enums.WIDTH_TILE

		offsetY := posY * 4 * enums.HEIGHT_TILE * enums.WIDTH
		offsetX := posX * enums.WIDTH_TILE

		startPix := i*enums.WIDTH*4 + offsetY + offsetX
		endPix := startPix + enums.WIDTH_TILE

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return Sprite{
		tileType:   tileType,
		PixelArray: pixel,
	}
}
