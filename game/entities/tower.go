package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
)

type Tower struct {
	x       int
	y       int
	removed bool
	level   level.Level
	sprite  render.Sprite
}

func NewTower(x int, y int) Entity {
	return &Tower{
		x: x,
		y: y,
	}
}

func (tower *Tower) Update() {

}
func (tower *Tower) Render() {

}
func (tower *Tower) GetX() int {
	return tower.x
}
func (tower *Tower) GetY() int {
	return tower.y
}
func (tower *Tower) Remove() {

}
func (tower *Tower) IsRemoved() bool {
	return false
}
