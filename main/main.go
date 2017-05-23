package main

import (
	"math/rand"
	"time"

	"github.com/LokiTheMango/jatdg/enums"
	"github.com/LokiTheMango/jatdg/graphics"
)

func main() {
	//TODO: INIT GAME
	graphics.InitWindowLoop("GAME", 160*4, 160*4, 160, 160, func(sharedWindow *graphics.Window) {
		startGame(sharedWindow)
	})

}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func startGame(window *graphics.Window) {

	tilemap, size := graphics.NewTileMap("C:\\Projects\\Go\\src\\github.com\\LokiTheMango\\jatdg\\resources\\tiles.jpg")
	tiles := createTileArray(tilemap.PixelArray)
	framebuffer := parseFrameBuffer(tiles, size)
	lastVBlankTime := time.Now()
	spent := time.Now().Sub(lastVBlankTime)
	spent++
	window.Mutex.Lock()
	copy(window.Pixel, framebuffer)
	window.RequestDraw()
	window.Mutex.Unlock()
	for {
		/*
			spent := time.Now().Sub(lastVBlankTime)
			toWait := 17*time.Millisecond - spent
			if toWait > time.Duration(0) {
				<-time.NewTimer(toWait).C
			}

			window.Mutex.Lock()
			newInput := game.Input {
				Keys: game.Keys {
					Sel:  window.CharIsDown('t'), Start: window.CharIsDown('y'),
					Up:   window.CharIsDown('w'), Down:  window.CharIsDown('s'),
					Left: window.CharIsDown('a'), Right: window.CharIsDown('d'),
					A:    window.CharIsDown('k'), B:     window.CharIsDown('j'),
				},
			}
			numDown := 'x'
			for r := '0'; r <= '9'; r++ {
				if window.CharIsDown(r) {
					numDown = r
					break
				}
			}
			window.Mutex.Unlock()*/
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
