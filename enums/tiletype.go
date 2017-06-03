package enums

type TileType int

const (
	VOID       = iota
	WALL       = iota
	TOWER      = iota
	ENEMY      = iota
	PROJECTILE = iota
	GOAL       = iota
)
