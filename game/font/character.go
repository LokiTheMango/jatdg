package font

import (
	"strings"

	"fmt"

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

func (char *Character) GetPixelArray() []byte {
	return char.PixelArray
}

func StringToCharArr(pixelArray []byte, str string) []Character {
	result := make([]Character, len(str))
	upperStr := []rune(strings.ToUpper(str))
	for i, runeV := range upperStr {
		fmt.Println(runeV)
		index := enums.GetIndexFromRune(runeV)
		posX, posY := int(index%10), int(index/10)
		pix := pickOutChar(pixelArray, index, posX, posY)
		result[i] = Character{
			charIndex:  index,
			PixelArray: pix,
			X:          posX,
			Y:          posY,
		}
	}
	return result
}

func pickOutChar(pixelArray []byte, charIndex enums.CharacterIndex, posX int, posY int) []byte {
	pixel := make([]byte, 16*4*16)
	for i := 0; i < 16; i++ {
		start := i * 16 * 4
		end := start + 16*4

		offsetY := posY * 16 * (enums.WIDTH / 2)
		offsetX := posX * 16 * 4

		startPix := i*(enums.WIDTH/2) + offsetY + offsetX
		endPix := startPix + 16*4

		copy(pixel[start:end], pixelArray[startPix:endPix])
	}
	return pixel
}
