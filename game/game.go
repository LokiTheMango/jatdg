package game

import (
	"fmt"
	"math"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/input"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

// Game Object
type Game struct {
	screen         Screen
	level          level.Level
	input          input.Keyboard
	DrawRequested  bool
	Towers         map[int]entities.Entity
	Spawns         map[int]entities.Entity
	Enemies        map[int]entities.Mob
	Projectiles    map[int]entities.Mob
	camera         entities.Mob
	time           int
	numEnemies     int
	numProjectiles int
}

//Constructor
func New() *Game {
	game := &Game{
		screen:         Screen{},
		input:          input.Keyboard{},
		DrawRequested:  false,
		time:           0,
		numEnemies:     0,
		numProjectiles: 0,
	}
	return game
}

func (game *Game) Init(filePath string) {
	game.screen = NewScreen(filePath + "tiles.png")
	level, towerTiles, spawnTiles := level.NewLevel(game.screen.SpriteSheet, filePath+"level.png")
	game.level = level
	game.screen.SetLevel(&game.level)
	game.camera = entities.NewCamera(&game.level, 1)
	game.createTowerEntities(towerTiles)
	game.createSpawnerEntities(spawnTiles)
	game.Enemies = make(map[int]entities.Mob)
	game.Projectiles = make(map[int]entities.Mob)
}

func (game *Game) createTowerEntities(tiles []*tiles.Tile) {
	game.Towers = make(map[int]entities.Entity)
	i := 0
	for j := 0; j < len(tiles); j++ {
		if tiles[j] == nil {
			continue
		}
		game.Towers[i] = entities.NewTower(tiles[j], i)
		i++
	}
}

func (game *Game) createSpawnerEntities(tiles []*tiles.Tile) {
	game.Spawns = make(map[int]entities.Entity)
	i := 0
	for j := 0; j < len(tiles); j++ {
		if tiles[j] == nil {
			continue
		}
		game.Spawns[i] = entities.NewEnemySpawn(tiles[j].X, tiles[j].Y, i)
		i++
	}
}

func (game *Game) Update() {
	game.time++
	game.spawnEnemies()
	for _, tower := range game.Towers {
		tower.Update()
	}
	for _, spawn := range game.Spawns {
		spawn.Update()
	}
	game.moveObjects()
	game.firingProjectiles()
	game.hitCheck()
	game.removeDeadObjects()
	game.clearScreen()
	game.render()
	for _, enemy := range game.Enemies {
		tile := *enemy.GetTile()
		game.screen.RenderMob(enemy.GetX(), enemy.GetY(), tile)
	}
	for _, projectile := range game.Projectiles {
		tile := *projectile.GetTile()
		game.screen.RenderMob(projectile.GetX(), projectile.GetY(), tile)
	}
}

func (game *Game) spawnEnemies() {
	if game.numEnemies < 10 {
		for _, spawn := range game.Spawns {
			if spawn.ReadyCheck() {
				fmt.Println("created Enemy")
				x := spawn.GetX()
				y := spawn.GetY()
				tile := game.level.CreateEnemy(x, y)
				game.Enemies[game.numEnemies] = entities.NewEnemy(tile, game.numEnemies, 10)
				game.numEnemies++
				spawn.Unready()
			}
		}
	}
}

func (game *Game) render() {
	x := game.camera.GetX()
	y := game.camera.GetY()
	game.screen.RenderLevel(x, y)
}

func (game *Game) clearScreen() {
	game.screen.ClearScreen()
}

func (game *Game) moveObjects() {
	for _, projectile := range game.Projectiles {
		projectile.Move(0, 0)
	}
	xOffset := 0
	yOffset := 0
	if game.input.Up {
		yOffset--
	}
	if game.input.Down {
		yOffset++
	}
	if game.input.Left {
		xOffset--
	}
	if game.input.Right {
		xOffset++
	}
	if xOffset != 0 || yOffset != 0 {
		game.camera.Move(xOffset, yOffset)
		game.moveWithCollisionCheck(xOffset, 0)
		game.moveWithCollisionCheck(0, yOffset)
	}
}

func (game *Game) moveWithCollisionCheck(xa int, ya int) {
	for _, enemy := range game.Enemies {
		x := int((float64(enemy.GetX()+xa) / 128))
		x2 := int((float64(enemy.GetX()+127+xa) / 128))
		y := (int((float64(enemy.GetY()+ya) / 32))) * game.level.Width
		y2 := (int((float64(enemy.GetY()+31+ya) / 32))) * game.level.Width
		if !game.collisionCheckForTile(x, x2, y, y2) {
			enemy.Move(xa<<2, ya)
		}
	}
}

func (game *Game) collisionCheckForTile(x1 int, x2 int, y1 int, y2 int) bool {
	collision := false
	collision = game.level.Tiles[x1+y1].TileProperties.IsSolid || collision
	collision = game.level.Tiles[x1+y2].TileProperties.IsSolid || collision
	collision = game.level.Tiles[x2+y1].TileProperties.IsSolid || collision
	collision = game.level.Tiles[x2+y2].TileProperties.IsSolid || collision
	return collision
}

func (game *Game) firingProjectiles() {
	for _, tower := range game.Towers {
		x := float64(tower.GetX()<<5) + enums.WIDTH_TILE/2
		y := float64(tower.GetY()<<5) + enums.HEIGHT_TILE/2
		for _, enemy := range game.Enemies {
			if tower.ReadyCheck() {
				xa := float64(enemy.GetX()/4 + enums.WIDTH_TILE/2)
				ya := float64(enemy.GetY() + enums.HEIGHT_TILE/2)
				xd := x - xa
				yd := y - ya
				check := math.Abs(xd) + math.Abs(yd)
				if check <= 128 {
					alpha := math.Atan2(yd, xd)
					fmt.Println("shooting")
					fmt.Println(alpha * (180 / math.Pi))
					game.shoot(tower.GetX(), tower.GetY(), alpha)
					tower.Unready()
				}
			}
		}
	}
}

func (game *Game) shoot(x int, y int, angle float64) {
	tile := game.level.CreateProjectile(x, y)
	game.Projectiles[game.numProjectiles] = entities.NewProjectile(tile, game.numProjectiles, angle, 1, 100)
	game.numProjectiles++
}

func (game *Game) hitCheck() {
	for _, projectile := range game.Projectiles {
		x := float64(projectile.GetX()<<5) + enums.WIDTH_TILE/2
		y := float64(projectile.GetY()<<5) + enums.HEIGHT_TILE/2
		for _, enemy := range game.Enemies {
			xa := float64(enemy.GetX()/4 + enums.WIDTH_TILE/2)
			ya := float64(enemy.GetY() + enums.HEIGHT_TILE/2)
			check := math.Abs(x-xa) <= 10 || math.Abs(y-ya) <= 10
			if check {
				projectile.Hit(1000)
				enemy.Hit(10)
			}
		}
	}
}

func (game *Game) removeDeadObjects() {
	for _, projectile := range game.Projectiles {
		if projectile.IsRemoved() {
			delete(game.Projectiles, projectile.GetIndex())
			game.numProjectiles--
		}
	}
	for _, enemy := range game.Enemies {
		if enemy.IsRemoved() {
			delete(game.Enemies, enemy.GetIndex())
			game.numEnemies--
		}
	}
}

func (game *Game) UpdateInput(newInput input.Keyboard) {
	game.input = newInput
}

func (game *Game) GetSpriteSheet() render.SpriteSheet {
	return game.screen.SpriteSheet
}

func (game *Game) GetPixelArray() []byte {
	return game.screen.PixelArray
}
