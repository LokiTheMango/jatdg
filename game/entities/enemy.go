package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Enemy struct {
	removed bool
	level   level.Level
	tile    *tiles.Tile
	X, Y    int
	hp      int
	index   int
	path    []pathing.Node
}

func NewEnemy(tile *tiles.Tile, index int, hp int) Mob {
	return &Enemy{
		tile:    tile,
		X:       tile.X << 7,
		Y:       tile.Y << 5,
		removed: false,
		index:   index,
	}
}

func (enemy *Enemy) GetIndex() int {
	return enemy.index
}
func (enemy *Enemy) SetPath(path []pathing.Node) {
	enemy.path = path
}
func (enemy *Enemy) Move(xa int, ya int) {
	if enemy.path != nil {
		if len(enemy.path) > 0 {
			index := len(enemy.path) - 1
			vec := enemy.path[index].Tile
			vx, vy := vec.GetXY()
			if enemy.X < vx<<7 {
				xa++
			}
			if enemy.X > vx<<7 {
				xa--
			}
			if enemy.Y < vy<<5 {
				ya++
			}
			if enemy.Y > vy<<5 {
				ya--
			}
		}
	}
	enemy.X += xa << 2
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
func (enemy *Enemy) Hit(dmg int) {
	enemy.hp -= dmg
	if enemy.hp <= 0 {
		enemy.Remove()
	}
}
func (enemy *Enemy) Remove() {
	enemy.removed = true
}
func (enemy *Enemy) IsRemoved() bool {
	return enemy.removed
}
