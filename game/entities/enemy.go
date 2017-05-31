package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Enemy struct {
	removed bool
	level   level.Level
	tile    *tiles.Tile
	X       int
	Y       int
}

func NewEnemy(tile *tiles.Tile) Mob {
	return &Enemy{
		tile: tile,
		X:    tile.X << 7,
		Y:    tile.Y << 5,
	}
}

func (enemy *Enemy) Move(xa int, ya int) {
	enemy.X += xa
	enemy.Y += ya
}
func (enemy *Enemy) GetX() int {
	return enemy.X
}
func (enemy *Enemy) GetY() int {
	return enemy.Y
}
func (enemy *Enemy) GetTile() *tiles.Tile {
	return enemy.tile
}
func (enemy *Enemy) Remove() {

}
func (enemy *Enemy) IsRemoved() bool {
	return false
}
