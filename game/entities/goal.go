package entities

import "github.com/LokiTheMango/jatdg/game/tiles"

type Goal struct {
	removed bool
	tile    *tiles.Tile
	index   int
}

func NewGoal(tile *tiles.Tile, index int) Entity {
	return &Goal{
		removed: false,
		tile:    tile,
		index:   index,
	}
}

func (goal *Goal) GetIndex() int {
	return goal.index
}
func (goal *Goal) Update() {

}
func (goal *Goal) ReadyCheck() bool {
	return false
}
func (goal *Goal) Unready() {

}
func (goal *Goal) GetX() int {
	return goal.tile.X
}
func (goal *Goal) GetY() int {
	return goal.tile.Y
}
func (goal *Goal) Remove() {

}
func (goal *Goal) IsRemoved() bool {
	return false
}
