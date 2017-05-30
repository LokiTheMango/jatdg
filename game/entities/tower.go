package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Tower struct {
	removed bool
	level   level.Level
	tile    *tiles.Tile
}

func NewTower(tile *tiles.Tile) Entity {
	return &Tower{
		tile: tile,
	}
}

func (tower *Tower) Update() {

}
func (tower *Tower) Render() {

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
