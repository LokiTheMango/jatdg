package tiles

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type EnemyTile struct {
	Tile Tile
}

func NewEnemyTile(posX int, posY int, pixelArray []byte) EnemyTile {
	//Tower Tile Sprite at 3, 0
	tile := Tile{
		TileType: enums.TileType(enums.ENEMY),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.ENEMY),
			3,
			0),
		X: posX,
		Y: posY,
		TileProperties: TileProperties{
			IsSolid: true,
		},
	}
	return EnemyTile{
		Tile: tile,
	}
}
