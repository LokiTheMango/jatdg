package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/LokiTheMango/jatdg/game"
	"github.com/LokiTheMango/jatdg/game/input"
	"github.com/LokiTheMango/jatdg/graphics"
)

func main() {
	//TODO: INIT GAME
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	fmt.Println([]rune("a")[0])
	exPath := path.Dir(dir)
	fmt.Println(exPath)
	game := game.New()
	graphics.InitWindowLoop("GAME", 640, 640, 320, 320, func(sharedWindow *graphics.Window) {
		startGame(sharedWindow, game, exPath)
	})

}

func startGame(window *graphics.Window, gameI *game.Game, filePath string) {
	//FOR RELEASE : gameI.Init(filePath + "/resources/tiles.jpg")
	gameI.Init("C:\\Projects\\go\\src\\github.com\\LokiTheMango\\jatdg\\resources\\")
	lastVBlankTime := time.Now()

	for {
		if time.Now().Sub(lastVBlankTime) > time.Millisecond*8 {
			gameI.DrawRequested = true
			gameI.Update()
			if window.StopDrawing {
				toWait := 8 * time.Millisecond
				<-time.NewTimer(toWait).C
				lastVBlankTime = time.Now()
			}
		}
		if gameI.DrawRequested && !window.StopDrawing {
			window.Mutex.Lock()
			copy(window.Pixel, gameI.GetPixelArray())
			window.RequestDraw()
			window.Mutex.Unlock()
			spent := time.Now().Sub(lastVBlankTime)
			toWait := 8*time.Millisecond - spent
			if toWait > time.Duration(0) {
				<-time.NewTimer(toWait).C
			}
			lastVBlankTime = time.Now()
			gameI.DrawRequested = false
		}
		window.Mutex.Lock()
		newInput := input.Keyboard{
			Up: window.CharIsDown('w'), Down: window.CharIsDown('s'),
			Left: window.CharIsDown('a'), Right: window.CharIsDown('d'),
			Enter: window.KeycodeIsDown(40),
		}
		window.Mutex.Unlock()
		gameI.UpdateInput(newInput)
	}
}
