package main

import (
	"time"

	"github.com/LokiTheMango/jatdg/game"
	"github.com/LokiTheMango/jatdg/graphics"
)

func main() {
	//TODO: INIT GAME
	game := game.New()
	graphics.InitWindowLoop("GAME", 160*4, 160*4, 160, 160, func(sharedWindow *graphics.Window) {
		startGame(sharedWindow, game)
	})

}

func startGame(window *graphics.Window, gameI *game.Game) {

	gameI.CreateTileMap("C:\\Projects\\Go\\src\\github.com\\LokiTheMango\\jatdg\\resources\\tiles.jpg")
	gameI.CreateTileArray()
	framebuffer := gameI.ParseFrameBuffer()
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
