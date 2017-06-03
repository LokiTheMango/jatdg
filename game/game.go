package game

import (
	"fmt"
	"math"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game/entities"
	"github.com/LokiTheMango/jatdg/game/input"
	"github.com/LokiTheMango/jatdg/game/level"
	"github.com/LokiTheMango/jatdg/game/pathing"
	"github.com/LokiTheMango/jatdg/game/render"
	"github.com/LokiTheMango/jatdg/game/tiles"
)

// Game Object
type Game struct {
	lives          int
	screen         Screen
	level          level.Level
	input          input.Keyboard
	DrawRequested  bool
	Goals          map[int]entities.Entity
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
		lives:          20,
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
	level, towerTiles, spawnTiles, goalTiles := level.NewLevel(game.screen.SpriteSheet, filePath+"level.png")
	game.level = level
	game.screen.SetLevel(&game.level)
	game.camera = entities.NewCamera(&game.level, 1)
	game.createGoalEntities(goalTiles)
	game.createTowerEntities(towerTiles)
	game.createSpawnerEntities(spawnTiles)
	game.Enemies = make(map[int]entities.Mob)
	game.Projectiles = make(map[int]entities.Mob)
}

func (game *Game) createGoalEntities(tiles []*tiles.Tile) {
	game.Goals = make(map[int]entities.Entity)
	i := 0
	for j := 0; j < len(tiles); j++ {
		if tiles[j] == nil {
			continue
		}
		game.Goals[i] = entities.NewGoal(tiles[j], i)
		i++
	}
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
	game.checkGoalHit()
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

func (game *Game) checkGoalHit() {
	for _, goal := range game.Goals {
		x := float64(goal.GetX()/4 + enums.HEIGHT_TILE/2)
		y := float64(goal.GetY() + enums.HEIGHT_TILE/2)
		for _, enemy := range game.Enemies {
			xa := float64(enemy.GetX()/4 + enums.HEIGHT_TILE/2)
			ya := float64(enemy.GetY() + enums.HEIGHT_TILE/2)
			//Hit Box == circle over center (radius 10px)
			check := math.Abs(x-xa) <= 10 && math.Abs(y-ya) <= 10
			if check {
				game.lives--
				fmt.Println("GOAL HIT !!! LIVE DOWN")
				fmt.Println(game.lives)
			}
		}
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
				enemy := entities.NewEnemy(tile, game.numEnemies, 10)
				game.updateEnemyPath(enemy)
				game.Enemies[game.numEnemies] = enemy
				game.numEnemies++
				spawn.Unready()
			}
		}
	}
}

func (game *Game) updateEnemyPath(enemy entities.Mob) {
	gx := game.Goals[0].GetX()
	gy := game.Goals[0].GetY()
	destination := pathing.NewVector2i(gx, gy)
	ex := (enemy.GetX() + enums.WIDTH_TILE/2) >> 7
	ey := (enemy.GetY() + enums.HEIGHT_TILE/2) >> 5
	start := pathing.NewVector2i(ex, ey)
	enemy.SetPath(game.level.FindPath(start, destination))
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
		game.moveWithCollisionCheck()
	}
}

func (game *Game) moveWithCollisionCheck() {
	update := false
	if game.time%60 == 0 {
		update = true
	}
	for _, enemy := range game.Enemies {
		if update {
			game.updateEnemyPath(enemy)
		}
		x := int((float64(enemy.GetX()) / 128))
		x2 := int((float64(enemy.GetX()+127) / 128))
		y := (int((float64(enemy.GetY()) / 32))) * game.level.Width
		y2 := (int((float64(enemy.GetY()+31) / 32))) * game.level.Width
		if !game.collisionCheckForTile(x, x2, y, y2) {
			enemy.Move(0, 0)
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
	game.Projectiles[game.numProjectiles] = entities.NewProjectile(tile, game.numProjectiles, angle, 4, 100)
	game.numProjectiles++
}

func (game *Game) hitCheck() {
	for _, projectile := range game.Projectiles {
		x := float64(projectile.GetX()/4) + enums.HEIGHT_TILE/2
		y := float64(projectile.GetY()) + enums.HEIGHT_TILE/2
		for _, enemy := range game.Enemies {
			xa := float64(enemy.GetX()/4 + enums.HEIGHT_TILE/2)
			ya := float64(enemy.GetY() + enums.HEIGHT_TILE/2)
			//Hit Box == circle over center (radius 10px)
			check := math.Abs(x-xa) <= 10 && math.Abs(y-ya) <= 10
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
