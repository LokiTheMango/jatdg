package render

import (
	"github.com/LokiTheMango/jatdg/enums"
)

type Sprite struct {
	tileType   enums.TileType
	posX       int
	posY       int
	PixelArray []byte
}

func NewSprite(pixelArray []byte, tileType enums.TileType, posX int, posY int) Sprite {
	pixel := make([]byte, enums.WIDTH_TILE*enums.HEIGHT_TILE)
	var SpritePosX int
	var SpritePosY int
	switch tileType {
	case enums.VOID:
		SpritePosX = 0
		SpritePosY = 0
	case enums.BULLET:
		SpritePosX = 1
		SpritePosY = 0
	case enums.TOWER:
		SpritePosX = 2
		SpritePosY = 0
	case enums.ENEMY:
		SpritePosX = 3
		SpritePosY = 0
	}
	for i := 0; i < enums.HEIGHT_TILE; i++ {
		start := i * enums.WIDTH_TILE
		end := start + enums.WIDTH_TILE

		offsetY := SpritePosY * 4 * enums.HEIGHT_TILE * enums.WIDTH
		offsetX := SpritePosX * enums.WIDTH_TILE

		startPix := i*enums.WIDTH*4 + offsetY + offsetX
		endPix := startPix + enums.WIDTH_TILE

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return Sprite{
		tileType:   tileType,
		posX:       posX,
		posY:       posY,
		PixelArray: pixel,
	}
}
