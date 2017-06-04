package ui

import (
	"github.com/LokiTheMango/jatdg/game/font"
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/render"
)

type UIManager struct {
	Panels    []UIPanel
	FontSheet []byte
	size      int
}

func NewUIManager(filePath string) UIManager {
	fontSheet, size := render.NewFontSheet(filePath)
	return UIManager{
		Panels:    []UIPanel{},
		FontSheet: fontSheet.PixelArray,
		size:      size,
	}
}

func (ui *UIManager) AddPanel(panel UIPanel) {
	ui.Panels = append(ui.Panels, panel)
}

func (ui *UIManager) AddNewUITextTo(panelIndex int, x int, y int, str string) {
	if &ui.Panels[panelIndex] != nil {
		pos := pathing.NewVector2i(x, y)
		offs := pathing.NewVector2i(0, 0)
		font := font.StringToCharArr(ui.FontSheet, str)
		comp := NewUIText(pos, offs, font, str)
		ui.Panels[panelIndex].Components = append(ui.Panels[panelIndex].Components, comp)
	}
}

func (ui *UIManager) UpdateOffset(x int, y int) {
	offs := pathing.NewVector2i(x, y)
	for _, panel := range ui.Panels {
		panel.offset = offs
		panel.Update()
	}
}
