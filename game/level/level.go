package level

import (
	"math/rand"
	"sort"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/pathing"
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

	NumGoals      int
	NumTowers     int
	NumEnemySpawn int
}

func NewLevel(spriteSheet render.SpriteSheet, filePath string) (Level, []*tiles.Tile, []*tiles.Tile, []*tiles.Tile) {
	rand.Seed(time.Now().UTC().UnixNano())
	levelSheet, width, height := render.NewLevelSheet(filePath)
	level := Level{
		spriteSheet:   spriteSheet,
		levelSheet:    levelSheet,
		Width:         width,
		Height:        height,
		NumGoals:      0,
		NumTowers:     0,
		NumEnemySpawn: 0,
	}
	towerTiles, spawnTiles, goalTiles := level.generateLevel()
	return level, towerTiles, spawnTiles, goalTiles
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (level *Level) generateLevel() ([]*tiles.Tile, []*tiles.Tile, []*tiles.Tile) {
	level.Tiles = make([]tiles.Tile, level.Height*level.Width)
	towerTiles := make([]*tiles.Tile, level.Height*level.Width)
	spawnTiles := make([]*tiles.Tile, level.Height*level.Width)
	goalTiles := make([]*tiles.Tile, level.Height*level.Width)
	for i := 0; i < level.Height; i++ {
		for j := 0; j < level.Width; j++ {
			//// random Tile generation for Tests
			//nextTile := randInt(0, 2)
			nextTile := 0
			isTower := false
			isSpawn := false
			isGoal := false
			indexHeight := i * level.Width * 4
			pix := level.levelSheet.PixelArray[indexHeight+j*4 : indexHeight+(j+1)*4]
			wall := []byte{0, 0, 0, 255}
			goal := []byte{0, 255, 0, 255}
			tower := []byte{0, 0, 255, 255}
			spawn := []byte{255, 0, 0, 255}
			if testEq(pix, wall) {
				nextTile = 1
			} else if testEq(pix, goal) {
				nextTile = 5
				level.NumGoals++
				isGoal = true
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
			if isGoal == true {
				goalTiles[level.NumGoals-1] = &tile
			}
			if isTower == true {
				towerTiles[level.NumTowers-1] = &tile
			}
			if isSpawn == true {
				spawnTiles[level.NumTowers-1] = &tile
			}
			level.Tiles[i*level.Width+j] = tile
		}
	}
	return towerTiles, spawnTiles, goalTiles
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

func (level *Level) FindPath(start pathing.Vector2i, goal pathing.Vector2i) []pathing.Node {
	openList := []pathing.Node{}
	closedList := []pathing.Node{}
	current := pathing.NewNode(start, nil, 0, start.GetDistanceTo(&goal))
	openList = append(openList, current)
	for len(openList) > 0 {
		sort.Sort(pathing.ByFCost(openList))
		current = openList[0]
		if current.Tile.Equals(&goal) {
			path := []pathing.Node{}
			for current.Parent != nil {
				path = append(path, current)
				current = *current.Parent
			}
			openList = openList[:0]
			closedList = closedList[:0]
			return path
		}
		openList = openList[1:]
		closedList = append(closedList, current)
		// NO NEED FOR DIAGONALS
		// DIAGONALS IMPLEMENTED BY CHECKING 9 FIELDS
		// xi=i%3-1 yi=i/3-1
		for i := 0; i < 9; i++ {
			if i == 4 {
				continue
			}
			x, y := current.Tile.GetXY()
			xi := (i % 3) - 1
			yi := (i / 3) - 1
			if (xi == -1 && yi == -1) || (xi == 1 && yi == 1) || (xi == -1 && yi == 1) || (xi == 1 && yi == -1) {
				xi, yi = 0, 0
			}
			at := level.Tiles[(x+xi)+(y+yi)*level.Width]
			if &at == nil || at.TileProperties.IsSolid {
				continue
			}
			a := pathing.NewVector2i(x+xi, y+yi)
			gCost := current.GCost + current.Tile.GetDistanceTo(&a)
			hCost := a.GetDistanceTo(&goal)
			parent := current
			node := pathing.NewNode(a, &parent, gCost, hCost)
			if vecInSlice(closedList, &a) && gCost >= node.GCost {
				continue
			}
			if !vecInSlice(openList, &a) || gCost < node.GCost {
				openList = append(openList, node)
			}
		}
	}
	closedList = closedList[:0]
	return nil
}

func vecInSlice(slice []pathing.Node, vec *pathing.Vector2i) bool {
	for _, n := range slice {
		if n.Tile.Equals(vec) {
			return true
		}
	}
	return false
}
