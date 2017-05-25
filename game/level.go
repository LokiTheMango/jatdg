package game

import (
	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/render"
)

//Level Object
type Level struct {
	width  int
	height int

	PixelArray  []byte
	spriteSheet render.SpriteSheet
	tileSize    int
	tiles       []entities.Tile

	xOffset int
	yOffset int
}

func NewLevel(spriteSheet render.SpriteSheet, width int, height int) Level {
	level := Level{
		spriteSheet: spriteSheet,
		width:       width,
		height:      height,
	}
	level.generateLevel()
	return level
}

func (level *Level) generateLevel() {
	level.tiles = make([]entities.Tile, level.height*level.width)
	for i := 0; i < level.height; i++ {
		for j := 0; j < level.width; j++ {
			//// random Tile generation for Tests
			nextTile := randInt(0, 2)
			tileType := enums.TileType(nextTile)
			level.tiles[i*level.width+j] = entities.NewTile(j, i, tileType, level.spriteSheet.PixelArray)
		}
	}
}

func (level *Level) renderTile(xp int, yp int, tile entities.Tile) {
	xp = xp << 7
	yp = yp << 5
	xp -= level.xOffset
	yp -= level.yOffset
	tilePixels := tile.GetPixelArray()
	for y := 0; y < enums.HEIGHT_TILE; y++ {
		ya := y + yp
		for x := 0; x < enums.WIDTH_TILE; x++ {
			xa := x + xp
			if xa < 0 || xa >= enums.WIDTH || ya < 0 || ya >= enums.HEIGHT {
				break
			}
			indexPix := xa + ya*enums.WIDTH
			indexTilePix := x + y*enums.WIDTH_TILE
			level.PixelArray[indexPix] = tilePixels[indexTilePix]
		}
	}
}

func (level *Level) InitScreen(size int) {
	/*
		for y := 0; y < enums.HEIGHT; y++ {
			yOff := y + yOffset
			if yOff < 0 || yOff >= enums.HEIGHT {
				continue
			}
			// PIXEL WIDTH TIME 4 (RGBA)
			for x := 0; x < enums.WIDTH; x++ {
				xOff := x + (xOffset << 2)
				if xOff < 0 || xOff >= enums.WIDTH {
					continue
				}
				tileIndex := ((yOff >> 5) & (level.height - 1)) + ((xOff >> 7) & (level.width - 1))
				tileArr := level.tiles[tileIndex].GetPixelArray()
				tileWidth := enums.WIDTH_TILE
				tileHeight := enums.HEIGHT_TILE
				maskTileWidth := tileWidth - 1
				maskTileHeight := tileHeight - 1
				index := ((xOff & maskTileWidth) % tileWidth) + ((yOff&maskTileHeight)%tileHeight)*tileWidth
				framebuffer[x+y*enums.WIDTH] = tileArr[index]
			}
		}
	*/
	framebuffer := make([]byte, size)
	level.PixelArray = framebuffer
}

func (level *Level) setOffset(xOffset int, yOffset int) {
	level.xOffset = xOffset << 2
	level.yOffset = yOffset
}

func (level *Level) RenderLevel(xScroll int, yScroll int) {
	level.setOffset(xScroll, yScroll)
	x0 := xScroll >> 5
	x1 := (xScroll + enums.WIDTH) >> 5
	y0 := yScroll >> 5
	y1 := (yScroll + enums.HEIGHT) >> 5

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			index := x + y*level.width
			tile := level.tiles[index]
			level.renderTile(x, y, tile)
		}
	}
}
