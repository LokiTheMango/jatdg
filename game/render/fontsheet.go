package render

import (
	"image"
	"image/draw"

	//import png
	_ "image/png"
	"os"
)

type FontSheet struct {
	width      int
	height     int
	filePath   string
	PixelArray []byte
	encoding   string
}

//Constructor for FontSheet
func NewFontSheet(filePath string) (FontSheet, int) {
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, str, err := image.Decode(reader)
	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)
	return FontSheet{
		width:      (rect.Max.X - rect.Min.X),
		height:     (rect.Max.Y - rect.Min.Y),
		filePath:   filePath,
		PixelArray: rgba.Pix,
		encoding:   str,
	}, len(rgba.Pix)
}
