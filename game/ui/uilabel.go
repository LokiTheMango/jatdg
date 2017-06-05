package ui

import (
	"github.com/LokiTheMango/jatdg/game/font"
	"github.com/LokiTheMango/jatdg/game/pathing"
)

type UILabel struct {
	font             []font.Character
	label            string
	position, offset pathing.Vector2i
	pixelArray       []byte
}

func NewUILabel(label string, font []font.Character, position, offset pathing.Vector2i) UIComponent {
	return &UILabel{
		label:    label,
		font:     font,
		position: position,
		offset:   offset,
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
	return label.label
}

func (label *UILabel) GetPixelArray() []byte {
	pix := make([]byte, 16*4*16*len(label.font))
	for i, char := range label.font {
		copy(pix[i*16*4*16:(i+1)*16*4*16], char.GetPixelArray())
	}
	return pix
}

func (label *UILabel) GetPosition() pathing.Vector2i {
	return label.position
}
