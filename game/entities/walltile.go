package entities

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type WallTile struct {
	Tile Tile
}

func NewWallTile(posX int, posY int, pixelArray []byte) WallTile {
	//Wall Tile Sprite at 1, 0
	tile := Tile{
		TileType: enums.TileType(enums.WALL),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.WALL),
			1,
			0),
		X: posX,
		Y: posY,
		tileProperties: TileProperties{
			isSolid: true,
		},
	}
	return WallTile{
		Tile: tile,
	}
}
