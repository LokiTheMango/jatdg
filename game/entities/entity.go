package entities

import "github.com/LokiTheMango/jatdg/game/tiles"

type Entity interface {
	Update()
	Render()
	GetX() int
	GetY() int
	Remove()
	IsRemoved() bool
}

type Mob interface {
	GetX() int
	GetY() int
	GetTile() *tiles.Tile
	Remove()
	IsRemoved() bool
	Move(xa int, ya int)
}
