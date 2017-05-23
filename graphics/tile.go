package graphics

import "github.com/LokiTheMango/jatdg/enums"

type Tile struct {
	tileType   enums.TileType
	posX       int
	posY       int
	pixelArray []byte
}

func NewTile(pixelArray []byte, tileType enums.TileType, posX int, posY int) Tile {
	pixel := make([]byte, enums.WIDTH_TILE*enums.HEIGHT_TILE)
	var tilePosX int
	var tilePosY int
	switch tileType {
	case enums.VOID:
		tilePosX = 0
		tilePosY = 0
	case enums.BULLET:
		tilePosX = 1
		tilePosY = 0
	case enums.TOWER:
		tilePosX = 2
		tilePosY = 0
	case enums.ENEMY:
		tilePosX = 3
		tilePosY = 0
	}
	for i := 0; i < enums.HEIGHT_TILE; i++ {
		start := i * enums.WIDTH_TILE
		end := start + enums.WIDTH_TILE

		offsetY := tilePosY * 4 * enums.HEIGHT_TILE * enums.WIDTH
		offsetX := tilePosX * enums.WIDTH_TILE

		startPix := i*enums.WIDTH*4 + offsetY + offsetX
		endPix := startPix + enums.WIDTH_TILE

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return Tile{
		tileType:   tileType,
		posX:       posX,
		posY:       posY,
		pixelArray: pixel,
	}
}

func (tile *Tile) GetPixelArray() []byte {
	return tile.pixelArray
}
