package entities

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
	Remove()
	IsRemoved() bool
	Move(xa int, ya int)
}
