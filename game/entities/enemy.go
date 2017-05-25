package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
)

type Enemy struct {
	x       int
	y       int
	removed bool
	level   level.Level
	sprite  render.Sprite
}

func NewEnemy(x int, y int, sprite render.Sprite) Mob {
	return &Enemy{
		x:      x,
		y:      y,
		sprite: sprite,
	}
}

func (enemy *Enemy) Move(xa int, ya int) {
	if !enemy.collision() {
		enemy.x += xa
		enemy.y += ya
	}
}
func (enemy *Enemy) Update() {

}
func (enemy *Enemy) Render() {

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

func (enemy *Enemy) collision() bool {
	return false
}
