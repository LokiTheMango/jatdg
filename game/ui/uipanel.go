package ui

import (
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/render"
)

type UIPanel struct {
	Components       []UIComponent
	position, offset pathing.Vector2i
	Sprite           render.Sprite
}

func NewUIPanel(position pathing.Vector2i) UIPanel {
	return UIPanel{
		Components: []UIComponent{},
		position:   position,
		offset:     pathing.NewVector2i(0, 0),
		Sprite:     render.NewColorSprite(60, 320, []byte{50, 50, 50, 255}),
	}
}

func (ui *UIPanel) Update() {
	for _, component := range ui.Components {
		component.SetOffset(ui.position)
	}
}

func (ui *UIPanel) AddComponent(comp UIComponent) {
	ui.Components = append(ui.Components, comp)
}

func (ui *UIPanel) GetPositionXY() (int, int) {
	return ui.position.GetXY()
}

func (ui *UIPanel) GetOffsetXY() (int, int) {
	return ui.offset.GetXY()
}
