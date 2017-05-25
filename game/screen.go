package game

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

type Screen struct {
	level           level.Level
	SpriteSheet     render.SpriteSheet
	SpriteSheetSize int
	PixelArray      []byte

	xOffset int
	yOffset int
}

func NewScreen(filePath string) Screen {
	screen := Screen{
		level:           level.Level{},
		SpriteSheet:     render.SpriteSheet{},
		SpriteSheetSize: 0,
	}
	screen.createSpriteSheet(filePath)
	screen.initScreen()
	return screen
}

func (screen *Screen) SetLevel(level level.Level) {
	screen.level = level
}

func (screen *Screen) initScreen() {
	screen.PixelArray = make([]byte, screen.SpriteSheetSize)
}

func (screen *Screen) createSpriteSheet(filePath string) {
	spriteSheet, size := render.NewSpriteSheet(filePath)
	screen.SpriteSheet = spriteSheet
	screen.SpriteSheetSize = size
}

func (screen *Screen) setOffset(xOffset int, yOffset int) {
	screen.xOffset = xOffset << 2
	screen.yOffset = yOffset
}

func (screen *Screen) RenderLevel(xScroll int, yScroll int) {
	screen.setOffset(xScroll, yScroll)
	x0 := xScroll >> 5
	x1 := (xScroll + enums.WIDTH + enums.WIDTH_TILE) >> 5
	y0 := yScroll >> 5
	y1 := (yScroll + enums.HEIGHT + enums.HEIGHT_TILE) >> 5

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			tileIndex := screen.getTileIndex(x, y)
			tile := screen.level.Tiles[0]
			if tileIndex == -1 {
				tile = tiles.GetVoidTile(x, y, screen.SpriteSheet.PixelArray)
			} else {
				tile = screen.level.Tiles[tileIndex]
			}
			screen.renderTile(x, y, tile)
		}
	}
}

func (screen *Screen) renderTile(xp int, yp int, tile tiles.Tile) {
	xp = xp << 7
	yp = yp << 5
	xp -= screen.xOffset
	yp -= screen.yOffset
	tilePixels := tile.GetPixelArray()
	for y := 0; y < enums.HEIGHT_TILE; y++ {
		ya := y + yp
		for x := 0; x < enums.WIDTH_TILE; x++ {
			xa := x + xp
			if xa < (-1*enums.WIDTH_TILE) || xa >= enums.WIDTH || ya < 0 || ya >= enums.HEIGHT {
				break
			}
			if xa < 0 {
				xa = 0
			}
			indexPix := xa + ya*enums.WIDTH
			indexTilePix := x + y*enums.WIDTH_TILE
			screen.PixelArray[indexPix] = tilePixels[indexTilePix]
		}
	}
}

func (screen *Screen) getTileIndex(x int, y int) int {
	if x < 0 || y < 0 || x >= enums.LEVEL_WIDTH || y >= enums.LEVEL_WIDTH {
		return -1
	}
	tileIndex := x + y*enums.LEVEL_WIDTH
	return tileIndex
}

func (screen *Screen) ClearScreen() {
	for i := 0; i < len(screen.PixelArray); i++ {
		screen.PixelArray[i] = 0
	}
}
