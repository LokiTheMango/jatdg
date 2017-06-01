package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Tower struct {
	removed bool
	level   level.Level
	tile    *tiles.Tile
	time    int
	ready   bool
}

func NewTower(tile *tiles.Tile) Entity {
	return &Tower{
		tile:  tile,
		time:  0,
		ready: false,
	}
}

func (tower *Tower) Update() {
	tower.time++
	if tower.time > 60 {
		tower.ready = true
		tower.time = 0
	}
}
func (tower *Tower) ReadyCheck() bool {
	return tower.ready
}
func (tower *Tower) GetX() int {
	return tower.tile.X
}
func (tower *Tower) GetY() int {
	return tower.tile.Y
}
func (tower *Tower) Remove() {

}
func (tower *Tower) IsRemoved() bool {
	return false
}
