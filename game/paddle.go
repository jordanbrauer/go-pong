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
func (paddle *Paddle) Update(elapsedTime float32) {
	if IsKeyPressed(sdl.SCANCODE_UP) && !paddle.IsTouchingTop() {
		paddle.Position.Y -= paddle.Velocity.Y * elapsedTime
	}

	if IsKeyPressed(sdl.SCANCODE_DOWN) && !paddle.IsTouchingBottom() {
		paddle.Position.Y += paddle.Velocity.Y * elapsedTime
	}

	if IsKeyPressed(sdl.SCANCODE_RIGHT) && !paddle.IsTouchingRight() {
		paddle.Position.X += paddle.Velocity.X * elapsedTime
	}

	if IsKeyPressed(sdl.SCANCODE_LEFT) && !paddle.IsTouchingLeft() {
		paddle.Position.X -= paddle.Velocity.X * elapsedTime
	}
}

// Goal will update the paddle player's score for the next render.
func (paddle *Paddle) Goal() {
	paddle.Score.Update()
}

func (paddle *Paddle) IsTouchingTop() bool {
	return 0.0 >= (paddle.Position.Y - (paddle.Height / 2.0))
}

func (paddle *Paddle) IsTouchingBottom() bool {
	return float32(WindowHeight) <= (paddle.Position.Y + (paddle.Height / 2.0))
}

func (paddle *Paddle) IsTouchingLeft() bool {
	return 0.0 >= (paddle.Position.X - (paddle.Width / 2.0))
}

func (paddle *Paddle) IsTouchingRight() bool {
	return float32(WindowWidth) <= (paddle.Position.X + (paddle.Width / 2.0))
}
