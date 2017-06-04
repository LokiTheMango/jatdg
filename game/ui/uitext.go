package ui

import (
	"github.com/LokiTheMango/jatdg/game/font"
	"github.com/LokiTheMango/jatdg/game/pathing"
)

type UIText struct {
	font             []font.Character
	label            string
	position, offset pathing.Vector2i
}

func NewUIText(position, offset pathing.Vector2i, font []font.Character, label string) UIComponent {
	return &UIText{
		font:     font,
		label:    label,
		offset:   offset,
		position: position,
	}
}

func (text *UIText) GetBackgroundColor() []byte {
	return nil
}

func (text *UIText) GetLabel() string {
	return text.label
}

func (text *UIText) GetOffset() pathing.Vector2i {
	return text.offset
}

func (text *UIText) SetOffset(offset pathing.Vector2i) {
	text.offset = offset
}

func (text *UIText) GetPixelArray() []byte {
	return nil
}

func (text *UIText) GetPosition() pathing.Vector2i {
	return text.position
}
