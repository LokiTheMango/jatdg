package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Camera struct {
	x, y    int
	removed bool
	level   *level.Level
	index   int
}

func NewCamera(level *level.Level, index int) Mob {
	return &Camera{
		level:   level,
		removed: false,
		index:   index,
	}
}

func (camera *Camera) GetIndex() int {
	return camera.index
}

func (camera *Camera) SetPath(path []pathing.Node) {

}

func (camera *Camera) Move(xa int, ya int) {
	camera.x += xa
	camera.y += ya
}
func (camera *Camera) GetX() int {
	return camera.x
}
func (camera *Camera) GetY() int {
	return camera.y
}
func (camera *Camera) GetTile() *tiles.Tile {
	return &tiles.Tile{}
}
func (camera *Camera) Hit(dmg int) {

}
func (camera *Camera) Remove() {
	camera.removed = true
}
func (camera *Camera) IsRemoved() bool {
	return camera.removed
}
