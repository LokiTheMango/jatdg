package ui

import (
	"github.com/LokiTheMango/jatdg/game/font"
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/render"
)

type UILabel struct {
	Font             []font.Character
	Sprite           render.Sprite
	label            string
	position, offset pathing.Vector2i
	pixelArray       []byte
}

func NewUILabel(label string, font []font.Character, sprite render.Sprite, position, offset pathing.Vector2i) UIComponent {
	return &UILabel{
		label:    label,
		Font:     font,
		Sprite:   sprite,
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
	if len(label.Font) != 0 {
		pix := make([]byte, 16*4*16*len(label.Font))
		for i, char := range label.Font {
			copy(pix[i*16*4*16:(i+1)*16*4*16], char.GetPixelArray())
		}
		return pix
	} else if &label.Sprite != nil {
		return label.Sprite.PixelArray
	}
	return nil
}

func (label *UILabel) GetPosition() pathing.Vector2i {
	return label.position
}
