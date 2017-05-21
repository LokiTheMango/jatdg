package graphics

import (
	"sync"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Window struct {
	Width  		int
	Height 		int
	Mutex  		sync.Mutex

	Pixel  		[]byte

	keyCodeArray	[256]bool
	keyCodeMap	map[key.Code]bool
	keyCharArray	[256]bool
	keyCharMap	map[rune]bool

	eventQueue	screen.EventDeque
	drawRequested	bool
}

type drawRequest struct {}

func (w *Window) CharIsDown(c rune) bool {
	if c < 256 {
		return w.keyCharArray[byte(c)]
	}
	return w.keyCharMap[c]
}

func (w *Window) KeycodeIsDown(c key.Code) bool {
	if c < 256 {
		return w.keyCodeArray[byte(c)]
	}
	return w.keyCodeMap[c]
}

func (w *Window) updateKeyboardState(e key.Event) {
	setVal := e.Direction == key.DirPress
	if setVal ||e.Direction == key.DirRelease {
		if e.Code < 256 {
			w.keyCodeArray[byte(e.Code)] = setVal
		}
		w.keyCodeMap[e.Code] = setVal
		if e.Rune < 256 {
			w.keyCharArray[byte(e.Rune)] = setVal
		}
		w.keyCharMap[e.Rune] = setVal
	}
}

//Draw request to window loop
func (w *Window) RequestDraw() {
	if !w.drawRequested {
		w.eventQueue.Send(drawRequest{})
		w.drawRequested = true
	}
}

// create window and start event loop
func InitWindowLoop(windowTitle string, windowWidth int, windowHeight int, frameWidth int, frameHeight int, updateLoop func(*Window)) {
	driver.Main(func (s screen.Screen) {

		win, err := s.NewWindow(&screen.NewWindowOptions{windowWidth, windowHeight, windowTitle})
		if err != nil {
			panic(err)
		}
		defer win.Release()

		buf, err := s.NewBuffer(image.Point{frameWidth, frameHeight})
		if err != nil {
			panic(err)
		}
		tex, err := s.NewTexture(image.Point{frameWidth, frameHeight})
		if err != nil {
			panic(err)
		}

		window := Window{
			Width: frameWidth,
			Height: frameHeight,
			Pixel: make([]byte, 4*frameWidth*frameHeight),
			eventQueue: win,
			keyCodeMap: map[key.Code]bool{},
			keyCharMap: map[rune]bool{},
		}

		go updateLoop(&window)

		szRect := buf.Bounds()
		needFullRepaint := true

		for {
			publish := false

			switch e := win.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

			case key.Event:
				window.Mutex.Lock()
				window.updateKeyboardState(e)
				window.Mutex.Unlock()

			case drawRequest:
				window.Mutex.Lock()
				copy(buf.RGBA().Pix, window.Pixel)
				tex.Upload(image.Point{0, 0}, buf, buf.Bounds())
				window.drawRequested = false
				window.Mutex.Unlock()
				publish = true

			case size.Event:
				szRect = e.Bounds()

			case paint.Event:
				needFullRepaint = true
				publish = true
			}

			if publish {
				scaleFacX := float64(szRect.Max.X) / float64(tex.Bounds().Max.X)
				scaleFacY := float64(szRect.Max.Y) / float64(tex.Bounds().Max.Y)
				scaleFac := scaleFacX
				if scaleFac < scaleFac {
					scaleFac = scaleFacY
				}

				scaleFac = float64(int(scaleFac))
				newWidth := int(scaleFac * float64(tex.Bounds().Max.X))
				centerX := float64(szRect.Max.X/2 - newWidth/2)
				src2dst := f64.Aff3 {
					float64(int(scaleFac)), 0, centerX,
					0, float64(int(scaleFac)), 0,
				}
				identTrans := f64.Aff3 {
					1, 0, 0,
					0, 1, 0,
				}

				if needFullRepaint {
					win.DrawUniform(identTrans, color.Black, szRect, screen.Src, nil)
					needFullRepaint = false
				}
				win.Draw(src2dst, tex, tex.Bounds(), screen.Src, nil)
				win.Publish()
			}
		}
	})
}