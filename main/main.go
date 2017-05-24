package main

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/game"
	"github.com/LokiTheMango/jatdg/graphics"
)

func main() {
	//TODO: INIT GAME
	game := game.New()
	rand.Seed(time.Now().UTC().UnixNano())
	graphics.InitWindowLoop("GAME", 160*4, 160*4, 160, 160, func(sharedWindow *graphics.Window) {
		startGame(sharedWindow, game)
	})

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func startGame(window *graphics.Window, gameI *game.Game) {

	tilemap, size := graphics.NewTileMap("C:\\Projects\\Go\\src\\github.com\\LokiTheMango\\jatdg\\resources\\tiles.jpg")
	tiles := createTileArray(tilemap.PixelArray)
	framebuffer := parseFrameBuffer(tiles, size)
	lastVBlankTime := time.Now()

	for {
		if time.Now().Sub(lastVBlankTime) > time.Millisecond*16 {
			gameI.DrawRequested = true
			gameI.Update()
		}
		if gameI.DrawRequested && !window.StopDrawing {
			window.Mutex.Lock()
			copy(window.Pixel, framebuffer)
			window.RequestDraw()
			window.Mutex.Unlock()
			spent := time.Now().Sub(lastVBlankTime)
			toWait := 17*time.Millisecond - spent
			if toWait > time.Duration(0) {
				<-time.NewTimer(toWait).C
			}
			lastVBlankTime = time.Now()
			gameI.DrawRequested = false
		}
		window.Mutex.Lock()
		newInput := game.Input{
			Up: window.CharIsDown('w'), Down: window.CharIsDown('s'),
			Left: window.CharIsDown('a'), Right: window.CharIsDown('d'),
			Enter: window.KeycodeIsDown(40),
		}
		window.Mutex.Unlock()
		gameI.UpdateInput(newInput)
	}
}

func createTileArray(arr []byte) []graphics.Tile {
	tiles := make([]graphics.Tile, 10*10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			nextTile := randInt(0, 4)
			tiles[i+j] = graphics.NewTile(arr, enums.TileType(nextTile), i, j)
		}
	}
	return tiles
}

func parseFrameBuffer(tiles []graphics.Tile, size int) []byte {
	framebuffer := make([]byte, size)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			arr := tiles[i+j].GetPixelArray()
			for k := 0; k < 16; k++ {
				copy(framebuffer[(k+i*16)*640+j*64:(k+i*16)*640+(j+1)*64], arr[k*64:(k+1)*64])
			}
		}
	}
	return framebuffer
}
