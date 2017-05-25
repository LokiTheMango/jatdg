package tiles

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type VoidTile struct {
	Tile Tile
}

func NewVoidTile(posX int, posY int, pixelArray []byte) VoidTile {
	//Void Tile Sprite at 0, 0
	tile := Tile{
		TileType: enums.TileType(enums.VOID),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.VOID),
			0,
			0),
		X: posX,
		Y: posY,
		tileProperties: TileProperties{
			isSolid: false,
		},
	}
	return VoidTile{
		Tile: tile,
	}
}
