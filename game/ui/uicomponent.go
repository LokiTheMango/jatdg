package ui

import (
	"github.com/LokiTheMango/jatdg/game/pathing"
)

type UIComponent interface {
	GetPixelArray() []byte
	GetPosition() pathing.Vector2i
	GetLabel() string
	GetBackgroundColor() []byte
	SetOffset(offset pathing.Vector2i)
	GetOffset() pathing.Vector2i
}
