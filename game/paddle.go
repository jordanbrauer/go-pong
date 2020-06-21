package game

import (
	"github.com/jordanbrauer/gogame/entity"
	"github.com/veandco/go-sdl2/sdl"
)

// Paddle is the object in which players will control to try and bat the ball
// back and forth on screen.
type Paddle struct {
	entity.Position
	entity.Dimensions
	entity.Velocity
	Colour entity.Colour
	Score  Score
}

// Draw will place the paddle on screen at it's current location based on the
// user's input
func (paddle *Paddle) Draw(pixels []byte) {
	var xOrigin = int(paddle.Position.X - (paddle.Width / 2))
	var yOrigin = int(paddle.Position.Y - (paddle.Height / 2))
	var height = int(paddle.Height)
	var width = int(paddle.Width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			Pixel(int32((xOrigin + x)), int32((yOrigin + y)), paddle.Colour, pixels)
		}
	}
}

// Update will receive input from the user and update the paddle's state to
// reflect it's new position on screen when it is drawn.
func (paddle *Paddle) Update(keyboard []uint8, elapsedTime float32) {

	var halfHeight = paddle.Height / 2
	var isTouchingTop bool = (paddle.Position.Y - halfHeight) <= 0
	var isTouchingBottom bool = (paddle.Position.Y + halfHeight) >= float32(WindowHeight)

	if 0 != keyboard[sdl.SCANCODE_UP] && !isTouchingTop {
		paddle.Position.Y -= paddle.Velocity.Y * elapsedTime
	}

	if 0 != keyboard[sdl.SCANCODE_DOWN] && !isTouchingBottom {
		paddle.Position.Y += paddle.Velocity.Y * elapsedTime
	}

	var halfWidth = paddle.Width / 2
	var isTouchingLeft bool = (paddle.Position.X - halfWidth) <= 0
	var isTouchingRight bool = (paddle.Position.X + halfWidth) >= float32(WindowWidth)

	if 0 != keyboard[sdl.SCANCODE_RIGHT] && !isTouchingRight {
		paddle.Position.X += paddle.Velocity.X * elapsedTime
	}

	if 0 != keyboard[sdl.SCANCODE_LEFT] && !isTouchingLeft {
		paddle.Position.X -= paddle.Velocity.X * elapsedTime
	}
}

// Goal will update the paddle player's score for the next render.
func (paddle *Paddle) Goal() {
	paddle.Score.Update()
}
