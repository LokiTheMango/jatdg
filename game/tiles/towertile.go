package tiles

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type TowerTile struct {
	Tile Tile
}

func NewTowerTile(posX int, posY int, pixelArray []byte) TowerTile {
	//Tower Tile Sprite at 2, 0
	tile := Tile{
		TileType: enums.TileType(enums.TOWER),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.TOWER),
			2,
			0),
		X: posX,
		Y: posY,
		TileProperties: TileProperties{
			IsSolid: true,
		},
	}
	return TowerTile{
		Tile: tile,
	}
}
