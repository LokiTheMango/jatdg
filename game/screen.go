package game

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
	"github.com/LokiTheMango/jatdg/game/ui"
)

type Screen struct {
	level           *level.Level
	SpriteSheet     render.SpriteSheet
	SpriteSheetSize int
	PixelArray      []byte

	xOffset int
	yOffset int
}

func NewScreen(filePath string) Screen {
	screen := Screen{
		level:           &level.Level{},
		SpriteSheet:     render.SpriteSheet{},
		SpriteSheetSize: 0,
	}
	screen.createSpriteSheet(filePath)
	screen.initScreen()
	return screen
}

func (screen *Screen) SetLevel(level *level.Level) {
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
			screen.RenderTile(x, y, tile)
		}
	}
}

func (screen *Screen) RenderTile(xp int, yp int, tile tiles.Tile) {
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

func (screen *Screen) DrawRect(xp int, yp int, width int, height int, color []byte, fixed bool) {
	xp = xp << 2
	width = width << 2
	if fixed {
		xp -= screen.xOffset
		yp -= screen.yOffset
	}
	pxIndex := 0
	index := 0
	for x := xp; x < xp+width; x += 4 {
		if x < 0 || x >= enums.WIDTH || yp >= enums.HEIGHT {
			continue
		}
		if yp > 0 {
			index = x + yp*enums.WIDTH
			for i := 0; i < 4; i++ {
				screen.PixelArray[index+i] = color[i]
			}
		}
		if yp+height >= enums.HEIGHT {
			continue
		}
		if yp+height > 0 {
			index = x + (yp+height)*enums.WIDTH
			for i := 0; i < 4; i++ {
				screen.PixelArray[index+i] = color[i]
			}
		}
		pxIndex++
	}
	pxIndex = 0
	for y := yp; y < yp+height; y++ {
		if xp >= enums.WIDTH || y < 0 || y >= enums.HEIGHT {
			continue
		}
		if xp > 0 {
			index = xp + y*enums.WIDTH
			for i := 0; i < 4; i++ {
				screen.PixelArray[index+i] = color[i]
			}
		}
		if xp+width >= enums.WIDTH {
			continue
		}
		if xp+width > 0 {
			index = (xp + width) + y*enums.WIDTH
			for i := 0; i < 4; i++ {
				screen.PixelArray[index+i] = color[i]
			}
		}
		pxIndex++
	}
}

func (screen *Screen) RenderMob(xp int, yp int, tile tiles.Tile) {
	xp -= screen.xOffset
	yp -= screen.yOffset
	tilePixels := tile.GetPixelArray()
	invisPix := []byte{255, 0, 255, 255}
	for y := 0; y < enums.HEIGHT_TILE; y++ {
		ya := y + yp
		for x := 0; x < enums.HEIGHT_TILE; x++ {
			xa := x*4 + xp
			if xa < (-1*enums.WIDTH_TILE) || xa >= enums.WIDTH || ya < 0 || ya >= enums.HEIGHT {
				break
			}
			if xa < 0 {
				xa = 0
			}
			indexPix := xa + ya*enums.WIDTH
			indexTilePix := x*4 + y*enums.WIDTH_TILE
			tilePixelCheck := tilePixels[indexTilePix : indexTilePix+4]
			if !testEq(tilePixelCheck, invisPix) {
				for i := 0; i < 4; i++ {
					screen.PixelArray[indexPix+i] = tilePixels[indexTilePix+i]
				}
			}
		}
	}
}

func (screen *Screen) getTileIndex(x int, y int) int {
	width := screen.level.Width
	height := screen.level.Height
	if x < 0 || y < 0 || x >= width || y >= height {
		return -1
	}
	tileIndex := x + y*width
	return tileIndex
}

func (screen *Screen) ClearScreen() {
	for i := 0; i < len(screen.PixelArray); i++ {
		screen.PixelArray[i] = 0
	}
}

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func (screen *Screen) RenderSprite(xp int, yp int, sprite render.Sprite, fixed bool) {
	xp = xp << 2
	if fixed {
		xp -= screen.xOffset
		yp -= screen.yOffset
	}
	spritePix := sprite.PixelArray
	height := sprite.Height
	width := sprite.Width
	for y := 0; y < height; y++ {
		ya := y + yp
		for x := 0; x < width*4; x++ {
			xa := x + xp
			if xa < (-1*width) || xa >= enums.WIDTH || ya < 0 || ya >= enums.HEIGHT {
				break
			}
			if xa < 0 {
				xa = 0
			}
			indexPix := xa + ya*enums.WIDTH
			indexSpritePix := x + y*width
			screen.PixelArray[indexPix] = spritePix[indexSpritePix]
		}
	}
}

func (screen *Screen) RenderUI(ui *ui.UIManager) {
	for _, panel := range ui.Panels {
		posX, posY := panel.GetPositionXY()
		screen.RenderSprite(posX, posY, panel.Sprite, false)
		for _, component := range panel.Components {
			component.GetLabel()
		}
	}
	screen.DrawRect(250, 50, 20, 20, []byte{255, 0, 255, 255}, false)
}
