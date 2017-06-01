package tiles

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
)

type Tile struct {
	TileType       enums.TileType
	sprite         render.Sprite
	X              int
	Y              int
	TileProperties TileProperties
}

type TileProperties struct {
	IsSolid bool
}

type SpriteProperties struct {
	TileType   enums.TileType
	pixelArray []byte
	posX       int
	posY       int
}

func newTileProperties(isSolid bool) TileProperties {
	return TileProperties{
		IsSolid: isSolid,
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
	case enums.TOWER:
		tile = NewTowerTile(posX, posY, pixelArray).Tile
	case enums.WALL:
		tile = NewWallTile(posX, posY, pixelArray).Tile
	case enums.ENEMY:
		tile = NewEnemyTile(posX, posY, pixelArray).Tile
	case enums.PROJECTILE:
		tile = NewProjectileTile(posX, posY, pixelArray).Tile
	default:
		tile = NewVoidTile(posX, posY, pixelArray).Tile
	}
	return tile
}

func GetVoidTile(posX int, posY int, pixelArray []byte) Tile {
	tile := NewVoidTile(posX, posY, pixelArray).Tile
	return tile
}

//Getter for PixelArray
func (tile *Tile) GetPixelArray() []byte {
	return tile.sprite.PixelArray
}
