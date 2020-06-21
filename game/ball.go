package game

import (
	"github.com/jordanbrauer/gogame/entity"
)

// Ball is a 2D game object that the user and/or computer player can bounce back
// and forth using their paddles.
type Ball struct {
	// The ball's coordinates in a 2D game space.
	entity.Position

	// The ball's dimensions which descrive it's shape in a 2D game space.
	entity.Dimensions

	// The ball's velocity to determine how fast it is currently moving in space.
	entity.Velocity

	// The ball's RGB colour represented as RGB bytes.
	Colour entity.Colour
}

// Draw will fill pixels with a colour in the given slice of pixels indicating
// the balls current location to the player.
func (ball *Ball) Draw(pixels []byte) {
	for y := -ball.Radius; y < ball.Radius; y++ {
		for x := -ball.Radius; x < ball.Radius; x++ {
			if ((x * x) + (y * y)) > (ball.Radius * ball.Radius) {
				continue
			}

			Pixel(
				int32((ball.Position.X + float32(x))),
				int32((ball.Position.Y + float32(y))),
				ball.Colour,
				pixels,
			)
		}
	}
}

// Update will change the position of the ball based on it's current position in
// the game world.
func (ball *Ball) Update(paddleLeft *Paddle, paddleRight *Paddle, elapsedTime float32) {
	// 1. check for left paddle collision physics
	if (ball.Position.X - ball.Radius) < (paddleLeft.Position.X + (paddleLeft.Width / 2)) {
		if ball.Position.Y > paddleLeft.Position.Y-paddleLeft.Height/2 && ball.Position.Y < paddleLeft.Position.Y+paddleLeft.Height/2 {
			ball.Velocity.X = -ball.Velocity.X // minimum translation vector
			ball.Position.X = paddleLeft.Position.X + (paddleLeft.Width / 2.0) + ball.Radius

			if ball.Position.Y <= paddleLeft.Position.Y-(paddleLeft.Height/4.0) {
				ball.Velocity.Y = -ball.Velocity.Y
			} else if ball.Position.Y >= paddleLeft.Position.Y+(paddleLeft.Height/4.0) {
				ball.Velocity.Y = -ball.Velocity.Y
			}
		}
	}

	var exitScreenRight bool = (ball.Position.X + ball.Radius) >= float32(WindowWidth)

	if exitScreenRight {
		paddleLeft.Goal()
	}

	// 2. check for right paddle collision physics
	if (ball.Position.X + ball.Radius) > (paddleRight.Position.X + (paddleRight.Width / 2)) {
		if ball.Position.Y > paddleRight.Position.Y-paddleRight.Height/2 && ball.Position.Y < paddleRight.Position.Y+paddleLeft.Height/2 {
			ball.Velocity.X = -ball.Velocity.X // minimum translation vector
			ball.Position.X = paddleRight.Position.X - (paddleRight.Width / 2.0) - ball.Radius

			if ball.Position.Y <= paddleRight.Position.Y-(paddleRight.Height/4.0) {
				ball.Velocity.Y = -ball.Velocity.Y
			} else if ball.Position.Y >= paddleRight.Position.Y+(paddleRight.Height/4.0) {
				ball.Velocity.Y = -ball.Velocity.Y
			}
		}
	}

	var exitScreenLeft bool = (int(ball.Position.X - ball.Radius)) <= 0

	if exitScreenLeft {
		paddleRight.Goal()
	}

	// 3. Update ball position and game state
	ball.Position.X += ball.Velocity.X * elapsedTime
	ball.Position.Y += ball.Velocity.Y * elapsedTime

	if int(ball.Position.Y-ball.Radius) <= 0 || ball.Position.Y+ball.Radius >= float32(WindowHeight) {
		ball.Velocity.Y = -ball.Velocity.Y
	}

	if exitScreenLeft || exitScreenRight {
		ball.Velocity = Velocity(-ball.Velocity.X, -ball.Velocity.Y)
		ball.Position = Centre()
		State = Waiting
	}
}
