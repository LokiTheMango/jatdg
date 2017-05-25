package entities

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
)

type Tile struct {
	TileType       enums.TileType
	sprite         render.Sprite
	X              int
	Y              int
	tileProperties TileProperties
}

type TileProperties struct {
	isSolid bool
}

type SpriteProperties struct {
	TileType   enums.TileType
	pixelArray []byte
	posX       int
	posY       int
}

func newTileProperties(isSolid bool) TileProperties {
	return TileProperties{
		isSolid: isSolid,
	}
}

func newSpriteProperties(tileType enums.TileType, pixelArray []byte, posX int, posY int) SpriteProperties {
	return SpriteProperties{
		TileType:   tileType,
		pixelArray: pixelArray,
		posX:       posX,
		posY:       posY,
	}
}

func NewTile(posX int, posY int, tileType enums.TileType, pixelArray []byte) Tile {
	tile := matchTileType(posX, posY, tileType, pixelArray)
	return tile
}

func matchTileType(posX int, posY int, tileType enums.TileType, pixelArray []byte) Tile {
	tile := Tile{}
	switch tileType {
	case enums.WALL:
		tile = NewWallTile(posX, posY, pixelArray).Tile
	default:
		tile = NewVoidTile(posX, posY, pixelArray).Tile
	}
	return tile
}

func (tile *Tile) Render(x int, y int) {

}

//Getter for PixelArray
func (tile *Tile) GetPixelArray() []byte {
	return tile.sprite.PixelArray
}
