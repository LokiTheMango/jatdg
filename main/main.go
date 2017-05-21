package main

import (
	"github.com/LokiTheMango/jatdg/graphics"
	"time"
)

func main() {
	//TODO: INIT GAME
	graphics.InitWindowLoop("GAME",160*4, 144*4, 160, 144, func(sharedWindow *graphics.Window) {
		startGame(sharedWindow)
	})

}

func startGame(window *graphics.Window) {

	lastVBlankTime := time.Now()

	for {
		window.Mutex.Lock()
		framebuffer := make([]byte, 160*140*4)
		for i:=0; i < len(framebuffer); i++ {
			if(i%3==0){
				framebuffer[i] = 0;
			} else {
				framebuffer[i] = 255;
			}
		}
		copy(window.Pixel, framebuffer)
		window.RequestDraw()
		window.Mutex.Unlock()
		spent := time.Now().Sub(lastVBlankTime)
		toWait := 17*time.Millisecond - spent
		if toWait > time.Duration(0) {
			<-time.NewTimer(toWait).C
		}

		/*window.Mutex.Lock()
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