package enums

type CharacterIndex int

const (
	A = iota
	B = iota
	C = iota
	D = iota
	E = iota
	F = iota
	G = iota
	H = iota
	I = iota
	J = iota
	K = iota
	L = iota
	M = iota
	N = iota
	O = iota
	P = iota
	Q = iota
	R = iota
	S = iota
	T = iota
	U = iota
	V = iota
	W = iota
	X = iota
	Y = iota
	Z = iota
)

func GetIndexFromRune(r rune) CharacterIndex {
	return CharacterIndex(r - 65)
}
