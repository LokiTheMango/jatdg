package game

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
)

// Game Object
type Game struct {
	gameMap       Map
	input         Input
	DrawRequested bool
	tiles         []render.Tile
	tileMap       render.TileMap
	tileMapSize   int
	PixelArray    []byte
	xOffset       int
	yOffset       int
}

//Constructor
func New() *Game {
	rand.Seed(time.Now().UTC().UnixNano())
	game := &Game{
		gameMap:       Map{},
		input:         Input{},
		DrawRequested: false,
		tiles:         make([]render.Tile, 0),
		tileMap:       render.TileMap{},
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

func (game *Game) SetTiles(tiles []render.Tile) {
	game.tiles = tiles
}
func (game *Game) GetTiles() []render.Tile {
	return game.tiles
}

func (game *Game) CreateTileMap(filePath string) {
	tileMap, size := render.NewTileMap(filePath)
	game.tileMap = tileMap
	game.tileMapSize = size
}
func (game *Game) GetTileMap() render.TileMap {
	return game.tileMap
}

func (game *Game) GetPixelArray() []byte {
	return game.tileMap.PixelArray
}

func (game *Game) CreateTileArray() {
	game.tiles = make([]render.Tile, 10*10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			nextTile := randInt(0, 4)
			game.tiles[i+j] = render.NewTile(game.tileMap.PixelArray, enums.TileType(nextTile), i, j)
		}
	}
}

func (game *Game) ParseFrameBuffer() {
	framebuffer := make([]byte, game.tileMapSize)
	for y := 0; y < enums.HEIGHT; y++ {
		yy := y + game.yOffset
		/*if yy < 0 || yy >= enums.HEIGHT {
			break
		}*/
		for x := 0; x < enums.WIDTH*4; x++ {
			xx := x + (game.xOffset << 2)
			/*if xx < 0 || xx >= enums.WIDTH*4 {
				break
			}*/
			tileIndex := ((yy >> 5) & 9) + ((xx >> 7) & 9)
			tileArr := game.tiles[tileIndex].GetPixelArray()
			index := (xx % 128) + (yy%32)*128
			framebuffer[x+y*enums.WIDTH*4] = tileArr[index]
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
