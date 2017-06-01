package entities

import (
	"github.com/LokiTheMango/jatdg/game/level"
)

type EnemySpawn struct {
	x       int
	y       int
	removed bool
	level   level.Level
	time    int
	ready   bool
}

func NewEnemySpawn(x int, y int) Entity {
	return &EnemySpawn{
		x:     x,
		y:     y,
		time:  0,
		ready: false,
	}
}

func (enemySpawn *EnemySpawn) Update() {
	enemySpawn.time++
	if enemySpawn.time > 1000 {
		enemySpawn.ready = true
		enemySpawn.time = 0
	}
}
func (enemySpawn *EnemySpawn) ReadyCheck() bool {
	return enemySpawn.ready
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
