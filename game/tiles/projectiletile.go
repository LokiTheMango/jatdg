package tiles

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type ProjectileTile struct {
	Tile Tile
}

func NewProjectileTile(posX int, posY int, pixelArray []byte) ProjectileTile {
	//Projectile Tile Sprite at 4, 0
	tile := Tile{
		TileType: enums.TileType(enums.PROJECTILE),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.PROJECTILE),
			4,
			0),
		X: posX,
		Y: posY,
		TileProperties: TileProperties{
			IsSolid: true,
		},
	}
	return ProjectileTile{
		Tile: tile,
	}
}
