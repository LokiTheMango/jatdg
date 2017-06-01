package entities

import (
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Projectile struct {
	removed bool
	tile    *tiles.Tile
	X       int
	Y       int
	Angle   float64
	xa      int
	ya      int
}

func NewProjectile(tile *tiles.Tile, angle float64) Mob {
	return &Projectile{
		tile:  tile,
		X:     tile.X << 7,
		Y:     tile.Y << 5,
		Angle: angle,
		xa:    0,
		ya:    0,
	}
}

func (Projectile *Projectile) Move(x int, y int) {
	/*if angle / 2 ==  {

	} else {

	}
	xd :=
	yd :=
	Projectile.X += xd
	Projectile.Y += yd*/
}
func (Projectile *Projectile) GetX() int {
	return Projectile.X
}
func (Projectile *Projectile) GetY() int {
	return Projectile.Y
}
func (Projectile *Projectile) GetTile() *tiles.Tile {
	return Projectile.tile
}
func (Projectile *Projectile) Remove() {

}
func (Projectile *Projectile) IsRemoved() bool {
	return false
}
