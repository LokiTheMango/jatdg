package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Enemy struct {
	x       int
	y       int
	removed bool
	level   level.Level
	tile    *tiles.Tile
}

func NewEnemy(x int, y int, tile *tiles.Tile) Mob {
	return &Enemy{
		x:    x,
		y:    y,
		tile: tile,
	}
}

func (enemy *Enemy) Move(xa int, ya int) {
	enemy.x += xa
	enemy.y += ya
}
func (enemy *Enemy) GetX() int {
	return enemy.x
}
func (enemy *Enemy) GetY() int {
	return enemy.y
}
func (enemy *Enemy) Remove() {

}
func (enemy *Enemy) IsRemoved() bool {
	return false
}
