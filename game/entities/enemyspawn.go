package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
)

type EnemySpawn struct {
	x       int
	y       int
	removed bool
	level   level.Level
	sprite  render.Sprite
}

func NewEnemySpawn(x int, y int) Entity {
	return &EnemySpawn{
		x: x,
		y: y,
	}
}

func (enemySpawn *EnemySpawn) Update() {

}
func (enemySpawn *EnemySpawn) Render() {

}
func (enemySpawn *EnemySpawn) GetX() int {
	return enemySpawn.x
}
func (enemySpawn *EnemySpawn) GetY() int {
	return enemySpawn.y
}
func (enemySpawn *EnemySpawn) Remove() {

}
func (enemySpawn *EnemySpawn) IsRemoved() bool {
	return false
}
