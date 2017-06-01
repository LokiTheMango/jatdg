package entities

import (
	"math"

	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Projectile struct {
	removed bool
	tile    *tiles.Tile
	X, Y    int
	Angle   float64
	xa, ya  float64
	speed   int
	hp      int
	index   int
}

func NewProjectile(tile *tiles.Tile, index int, angle float64, speed int, hp int) Mob {
	return &Projectile{
		tile:  tile,
		X:     tile.X << 7,
		Y:     tile.Y << 5,
		Angle: angle,
		speed: speed,
		xa:    float64(speed) * math.Cos(angle),
		ya:    float64(speed) * math.Sin(angle),
		hp:    hp,
		index: index,
	}
}

func (projectile *Projectile) GetIndex() int {
	return projectile.index
}
func (projectile *Projectile) Move(x int, y int) {
	projectile.X -= int(projectile.xa) << 2
	projectile.Y -= int(projectile.ya)
	projectile.Hit(1)
}
func (projectile *Projectile) GetX() int {
	return projectile.X
}
func (projectile *Projectile) GetY() int {
	return projectile.Y
}
func (projectile *Projectile) GetTile() *tiles.Tile {
	return projectile.tile
}
func (projectile *Projectile) Hit(dmg int) {
	projectile.hp -= dmg
	if projectile.hp <= 0 {
		projectile.Remove()
	}
}
func (projectile *Projectile) Remove() {
	projectile.removed = true
}
func (projectile *Projectile) IsRemoved() bool {
	return projectile.removed
}
