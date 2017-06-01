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
}

func NewProjectile(tile *tiles.Tile, angle float64, speed int) Mob {
	return &Projectile{
		tile:  tile,
		X:     tile.X << 7,
		Y:     tile.Y << 5,
		Angle: angle,
		speed: speed,
		xa:    float64(speed) * math.Cos(angle),
		ya:    float64(speed) * math.Sin(angle),
	}
}

func (projectile *Projectile) Move(x int, y int) {
	projectile.X -= int(projectile.xa) << 2
	projectile.Y -= int(projectile.ya)
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
func (projectile *Projectile) Remove() {

}
func (projectile *Projectile) IsRemoved() bool {
	return false
}
