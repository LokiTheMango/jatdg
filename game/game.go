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
	}
	return game
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (game *Game) Update() {

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

func (game *Game) ParseFrameBuffer() []byte {
	framebuffer := make([]byte, game.tileMapSize)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			arr := game.tiles[i+j].GetPixelArray()
			for k := 0; k < 16; k++ {
				copy(framebuffer[(k+i*16)*640+j*64:(k+i*16)*640+(j+1)*64], arr[k*64:(k+1)*64])
			}
		}
	}
	return framebuffer
}
