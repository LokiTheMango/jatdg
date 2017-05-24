package game

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
)

// Game Object
type Game struct {
	gameMap         Map
	input           Input
	DrawRequested   bool
	Sprites         []render.Sprite
	SpriteSheet     render.SpriteSheet
	SpriteSheetSize int
	PixelArray      []byte
	xOffset         int
	yOffset         int
}

//Constructor
func New() *Game {
	rand.Seed(time.Now().UTC().UnixNano())
	game := &Game{
		gameMap:       Map{},
		input:         Input{},
		DrawRequested: false,
		Sprites:       make([]render.Sprite, 0),
		SpriteSheet:   render.SpriteSheet{},
		xOffset:       0,
		yOffset:       0,
	}
	return game
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (game *Game) Update() {
	game.CheckInputForOffsets()
	game.ParseFrameBuffer()
}

func (game *Game) UpdateInput(newInput Input) {
	game.input = newInput
}

func (game *Game) SetSprites(Sprites []render.Sprite) {
	game.Sprites = Sprites
}
func (game *Game) GetSprites() []render.Sprite {
	return game.Sprites
}

func (game *Game) CreateSpriteSheet(filePath string) {
	spriteSheet, size := render.NewSpriteSheet(filePath)
	game.SpriteSheet = spriteSheet
	game.SpriteSheetSize = size
}
func (game *Game) GetSpriteMap() render.SpriteSheet {
	return game.SpriteSheet
}

func (game *Game) GetPixelArray() []byte {
	return game.SpriteSheet.PixelArray
}

func (game *Game) CreateSpriteArray() {
	game.Sprites = make([]render.Sprite, 10*10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			nextSprite := randInt(0, 4)
			game.Sprites[i+j] = render.NewSprite(game.SpriteSheet.PixelArray, enums.TileType(nextSprite), i, j)
		}
	}
}

func (game *Game) ParseFrameBuffer() {
	framebuffer := make([]byte, game.SpriteSheetSize)
	for y := 0; y < enums.HEIGHT; y++ {
		yy := y + game.yOffset
		if yy < 0 || yy >= enums.HEIGHT {
			continue
		}
		for x := 0; x < enums.WIDTH*4; x++ {
			xx := x + (game.xOffset << 2)
			if xx < 0 || xx >= enums.WIDTH*4 {
				continue
			}
			SpriteIndex := ((yy >> 5) & 9) + ((xx >> 7) & 9)
			SpriteArr := game.Sprites[SpriteIndex].GetPixelArray()
			index := ((xx & 127) % 128) + ((yy&31)%32)*128
			if index < 0 {
				index *= -1
			}
			framebuffer[x+y*enums.WIDTH*4] = SpriteArr[index]
		}
	}
	game.PixelArray = framebuffer
}

func (game *Game) CheckInputForOffsets() {
	if game.input.Up {
		game.yOffset++
	}
	if game.input.Down {
		game.yOffset--
	}
	if game.input.Left {
		game.xOffset++
	}
	if game.input.Right {
		game.xOffset--
	}
}
