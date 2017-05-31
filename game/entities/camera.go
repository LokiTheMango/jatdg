package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Camera struct {
	x       int
	y       int
	removed bool
	level   *level.Level
}

func NewCamera(level *level.Level) Mob {
	return &Camera{
		level: level,
	}
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
func (camera *Camera) Remove() {

}
func (camera *Camera) IsRemoved() bool {
	return false
}
