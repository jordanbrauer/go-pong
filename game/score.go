package game

import (
	"github.com/jordanbrauer/gogame/entity"
)

// Score is a representation of the position, size, colour, and value of a
// player's score in the game.
type Score struct {
	Number, Scale int
	entity.Position
	Colour entity.Colour
}

// A multi-dimensional array with each sub array containing a series of bytes
// that if arranged in a 3 ⨉ 5 matrix can be used to draw a number in a series
// of pixels
var numbers = [10][15]byte{
	{1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}, // 0
	{1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1}, // 1
	{1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, // 2
	{1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1}, // 3
	{1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1}, // 4
	{1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1}, // 5
	{0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1}, // 6
	{1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0}, // 7
	{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1}, // 8
	{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0}, // 9
}

// Draw the score at it's given location
func (score *Score) Draw(pixels []byte) {
	var xOrigin = int(score.Position.X) - ((score.Scale * 3) / 2)
	var yOrigin = int(score.Position.Y) - ((score.Scale * 5) / 2)
	var number = numbers[score.Number]

	for index, pixel := range number {
		if 1 == pixel {
			for y := yOrigin; y < (yOrigin + score.Scale); y++ {
				for x := xOrigin; x < (xOrigin + score.Scale); x++ {
					Pixel(int32(x), int32(y), score.Colour, pixels)
				}
			}
		}

		xOrigin += score.Scale

		if 0 == ((index + 1) % 3) {
			yOrigin += score.Scale
			xOrigin -= score.Scale * 3
		}
	}
}

// Update the score. If the score were to exceed 9, then it will roll over to 0
// to prevent out of range index errors. This is due to only being able to
// render single digits 0 through 9.
func (score *Score) Update() {
	if 9 == score.Number {
		score.Number = 0
	} else {
		score.Number++
	}
}
