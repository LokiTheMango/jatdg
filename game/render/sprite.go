package render

import (
	"github.com/LokiTheMango/jatdg/enums"
)

type Sprite struct {
	tileType   enums.TileType
	PixelArray []byte
	Width      int
	Height     int
}

func NewSprite(pixelArray []byte, tileType enums.TileType, posX int, posY int) Sprite {
	pixel := make([]byte, enums.WIDTH_TILE*enums.HEIGHT_TILE)
	for i := 0; i < enums.HEIGHT_TILE; i++ {
		start := i * enums.WIDTH_TILE
		end := start + enums.WIDTH_TILE

		offsetY := posY * enums.HEIGHT_TILE * enums.WIDTH
		offsetX := posX * enums.WIDTH_TILE

		startPix := i*enums.WIDTH + offsetY + offsetX
		endPix := startPix + enums.WIDTH_TILE

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return Sprite{
		tileType:   tileType,
		PixelArray: pixel,
	}
}

func NewColorSprite(width int, height int, color []byte) Sprite {
	pixel := make([]byte, width*4*height)
	for i := 0; i < height; i++ {
		for j := 0; j < width*4; j += 4 {
			start := j + i*width
			end := start + 4
			copy(pixel[start:end], color)
		}
	}
	return Sprite{
		PixelArray: pixel,
		Width:      width,
		Height:     height,
	}
}
