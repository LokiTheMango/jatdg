package tiles

import "github.com/LokiTheMango/jatdg/game/render"
import "github.com/LokiTheMango/jatdg/enums"

type GoalTile struct {
	Tile Tile
}

func NewGoalTile(posX int, posY int, pixelArray []byte) GoalTile {
	//Goal Tile Sprite at 5, 0
	tile := Tile{
		TileType: enums.TileType(enums.GOAL),
		sprite: render.NewSprite(
			pixelArray,
			enums.TileType(enums.GOAL),
			5,
			0),
		X: posX,
		Y: posY,
		TileProperties: TileProperties{
			IsSolid: false,
		},
	}
	return GoalTile{
		Tile: tile,
	}
}
