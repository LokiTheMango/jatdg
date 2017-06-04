package font

import (
	"strings"

	"github.com/LokiTheMango/jatdg/enums"
)

type Character struct {
	charIndex  enums.CharacterIndex
	PixelArray []byte
	X, Y       int
}

func NewCharacter(pixelArray []byte, charIndex enums.CharacterIndex) Character {
	posX, posY := int(charIndex%10), int(charIndex/10)
	pixel := pickOutChar(pixelArray, charIndex, posX, posY)
	return Character{
		charIndex:  charIndex,
		PixelArray: pixel,
		X:          posX,
		Y:          posY,
	}
}

func StringToCharArr(pixelArray []byte, str string) []Character {
	result := make([]Character, len(str))
	upperStr := []rune(strings.ToUpper(str))
	for i, runeV := range upperStr {
		index := enums.GetIndexFromRune(runeV)
		posX, posY := int(index%10), int(index/10)
		pixelArray := pickOutChar(pixelArray, index, posX, posY)
		result[i] = Character{
			charIndex:  index,
			PixelArray: pixelArray,
			X:          posX,
			Y:          posY,
		}
	}
	return result
}

func pickOutChar(pixelArray []byte, charIndex enums.CharacterIndex, posX int, posY int) []byte {
	pixel := make([]byte, enums.WIDTH_TILE*enums.HEIGHT_TILE)
	for i := 0; i < enums.HEIGHT_TILE; i++ {
		start := i * enums.WIDTH_TILE
		end := start + enums.WIDTH_TILE

		offsetY := posY * enums.HEIGHT_TILE * enums.WIDTH
		offsetX := posX * enums.WIDTH_TILE

		startPix := i*enums.WIDTH + offsetY + offsetX
		endPix := startPix + enums.WIDTH_TILE

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return pixel
}
