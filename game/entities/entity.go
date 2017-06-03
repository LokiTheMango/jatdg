package entities

import (
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Entity interface {
	GetIndex() int
	Update()
	ReadyCheck() bool
	Unready()
	GetX() int
	GetY() int
	Remove()
	IsRemoved() bool
}

type Mob interface {
	GetIndex() int
	SetPath(path []pathing.Node)
	GetX() int
	GetY() int
	GetTile() *tiles.Tile
	Remove()
	IsRemoved() bool
	Move(xa int, ya int)
	Hit(dmg int)
}
