package entities

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
)

type Tile struct {
	TileType enums.TileType
	sprite   render.Sprite
}

func NewTile(pixelArray []byte, tileType enums.TileType, posX int, posY int) Tile {
	return Tile{
		TileType: tileType,
		sprite:   render.NewSprite(pixelArray, tileType, posX, posY),
	}
}

//Getter for PixelArray
func (tile *Tile) GetPixelArray() []byte {
	return tile.sprite.PixelArray
}
