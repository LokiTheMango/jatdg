package pathing

import "math"

type Vector2i struct {
	x, y int
}

func NewVector2i(x int, y int) Vector2i {
	return Vector2i{
		x: x,
		y: y,
	}
}

func (vec *Vector2i) Set(x int, y int) {
	vec.x = x
	vec.y = y
}

func (vec *Vector2i) GetXY() (int, int) {
	return vec.x, vec.y
}

func (vec *Vector2i) Add(vec2 *Vector2i) {
	x2, y2 := vec2.GetXY()
	vec.x += x2
	vec.y += y2
}
func (vec *Vector2i) Sub(vec2 *Vector2i) {
	x2, y2 := vec2.GetXY()
	vec.x -= x2
	vec.y -= y2
}

func (vec *Vector2i) GetDistanceTo(vec2 *Vector2i) float64 {
	x2, y2 := vec2.GetXY()
	dx := float64(vec.x - x2)
	dy := float64(vec.y - y2)
	return math.Sqrt(dx*dx + dy*dy)
}
func (vec *Vector2i) Equals(vec2 *Vector2i) bool {
	x2, y2 := vec2.GetXY()
	if vec.x == x2 && vec.y == y2 {
		return true
	}
	return false
}
