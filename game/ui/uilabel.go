package ui

import (
	"github.com/LokiTheMango/jatdg/game/pathing"
)

type UILabel struct {
	text             string
	position, offset pathing.Vector2i
	pixelArray       []byte
}

func NewUILabel(pixelArray []byte, text string, position pathing.Vector2i) UIComponent {
	return &UILabel{
		text:       text,
		position:   position,
		pixelArray: pixelArray,
	}
}

func (label *UILabel) SetOffset(offset pathing.Vector2i) {
	label.offset = offset
}

func (label *UILabel) GetOffset() pathing.Vector2i {
	return label.offset
}

func (label *UILabel) GetBackgroundColor() []byte {
	return []byte{0, 0, 0, 255}
}

func (label *UILabel) GetLabel() string {
	return label.text
}

func (label *UILabel) GetPixelArray() []byte {
	return nil
}

func (label *UILabel) GetPosition() pathing.Vector2i {
	return label.position
}
