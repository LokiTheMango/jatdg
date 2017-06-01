package level

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

//Level Object
type Level struct {
	Width  int
	Height int

	spriteSheet render.SpriteSheet
	levelSheet  render.LevelSheet
	Tiles       []tiles.Tile

	NumTowers     int
	NumEnemySpawn int
}

func NewLevel(spriteSheet render.SpriteSheet, filePath string) (Level, []*tiles.Tile, []*tiles.Tile) {
	rand.Seed(time.Now().UTC().UnixNano())
	levelSheet, width, height := render.NewLevelSheet(filePath)
	level := Level{
		spriteSheet:   spriteSheet,
		levelSheet:    levelSheet,
		Width:         width,
		Height:        height,
		NumTowers:     0,
		NumEnemySpawn: 0,
	}
	towerTiles, spawnTiles := level.generateLevel()
	return level, towerTiles, spawnTiles
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (level *Level) generateLevel() ([]*tiles.Tile, []*tiles.Tile) {
	level.Tiles = make([]tiles.Tile, level.Height*level.Width)
	towerTiles := make([]*tiles.Tile, level.Height*level.Width)
	spawnTiles := make([]*tiles.Tile, level.Height*level.Width)
	for i := 0; i < level.Height; i++ {
		for j := 0; j < level.Width; j++ {
			//// random Tile generation for Tests
			//nextTile := randInt(0, 2)
			nextTile := 0
			isTower := false
			isSpawn := false
			indexHeight := i * level.Width * 4
			pix := level.levelSheet.PixelArray[indexHeight+j*4 : indexHeight+(j+1)*4]
			wall := []byte{0, 0, 0, 255}
			tower := []byte{0, 0, 255, 255}
			spawn := []byte{255, 0, 0, 255}
			if testEq(pix, wall) {
				nextTile = 1
			} else if testEq(pix, tower) {
				nextTile = 2
				level.NumTowers++
				isTower = true
			} else if testEq(pix, spawn) {
				nextTile = 0
				level.NumEnemySpawn++
				isSpawn = true
			}
			tileType := enums.TileType(nextTile)
			tile := tiles.NewTile(j, i, tileType, level.spriteSheet.PixelArray)
			if isTower == true {
				towerTiles[level.NumTowers-1] = &tile
			}
			if isSpawn == true {
				spawnTiles[level.NumTowers-1] = &tile
			}
			level.Tiles[i*level.Width+j] = tile
		}
	}
	return towerTiles, spawnTiles
}

func (level *Level) CreateEnemy(x int, y int) *tiles.Tile {
	tile := tiles.NewTile(x, y, enums.TileType(enums.ENEMY), level.spriteSheet.PixelArray)
	return &tile
}

func (level *Level) CreateProjectile(x int, y int) *tiles.Tile {
	tile := tiles.NewTile(x, y, enums.TileType(enums.PROJECTILE), level.spriteSheet.PixelArray)
	return &tile
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
