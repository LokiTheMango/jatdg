package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Enemy struct {
	removed bool
	level   level.Level
	tile    *tiles.Tile
}

func NewEnemy(tile *tiles.Tile) Mob {
	return &Enemy{
		tile: tile,
	}
}

func (enemy *Enemy) Move(xa int, ya int) {
	enemy.tile.X += xa
	enemy.tile.X += ya
}
func (enemy *Enemy) GetX() int {
	return enemy.tile.X
}
func (enemy *Enemy) GetY() int {
	return enemy.tile.Y
}
func (enemy *Enemy) GetTile() *tiles.Tile {
	return enemy.tile
}
func (enemy *Enemy) Remove() {

}
func (enemy *Enemy) IsRemoved() bool {
	return false
}
