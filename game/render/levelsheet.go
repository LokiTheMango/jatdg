package render

import (
	"image"
	"image/draw"

	//import png
	_ "image/png"
	"os"
)

type LevelSheet struct {
	width      int
	height     int
	filePath   string
	PixelArray []byte
	encoding   string
}

//Constructor for LevelSheet
func NewLevelSheet(filePath string) (LevelSheet, int, int) {
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, str, err := image.Decode(reader)
	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)
	width := rect.Max.X - rect.Min.X
	height := rect.Max.Y - rect.Min.Y
	return LevelSheet{
		width:      width,
		height:     height,
		filePath:   filePath,
		PixelArray: rgba.Pix,
		encoding:   str,
	}, width, height
}
