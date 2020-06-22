package game

import (
	"github.com/jordanbrauer/gogame/entity"
	"github.com/veandco/go-sdl2/sdl"
)

// Object is a simplate interface the depicts items in the game which can be
// drawn to screen as something the player can see.
type Object interface {
	Draw(pixels []byte)
}

// Mode can be one of many game states.
type Mode int

const (
	// Waiting game state mode represents the game paused and waiting for user
	// input to begin playing.
	Waiting Mode = iota

	// Playing game state mode represents the game being played and responding
	// to user input commands for the padddle.
	Playing
)

const (
	// BallRadius determines the total size of the pong ball.
	BallRadius float32 = 10

	// BallVelocity determines how fast the game ball travels when it is first
	// spawned,
	BallVelocity float32 = 400

	// MaxObjects is the absolute limit of Object structs that the game will
	// draw on screen.
	MaxObjects int = 5

	// PaddleLength is how tall/long the player paddles are.
	PaddleLength float32 = 75

	// PaddleWidth determines how wide/fat the player paddles are.
	PaddleWidth float32 = 10

	// PaddleVelocity determines how fast the player paddles move in relation to
	// the hardware and framerate over time.
	PaddleVelocity float32 = 400

	// ScoreBoardOffsetX is the percentage offset at which the scoreboards are
	// placed from the center of the screen on the X axis.
	ScoreBoardOffsetX float32 = 0.50

	// ScoreBoardOffsetY is the static offset at which the scoreboards are
	// placed from the top of the screen on the Y axis.
	ScoreBoardOffsetY float32 = 50

	// ScoreSize is the "font size" in which the numbers are rendered on screen.
	ScoreSize int = 10

	// WindowHeight is the default/assigned height of the game window, in pixels.
	WindowHeight int32 = 600

	// WindowWidth is the default/assigned width of the game window, in pixels.
	WindowWidth int32 = 800
)

// State is the current state of the game.
var State = Waiting
var keyboard = sdl.GetKeyboardState()

// Abort handles an error by checking for a nil value and panicing otherwise.
func Abort(caught error) {
	if caught != nil {
		panic(caught)
	}
}

// Lerp is a linear interpolation implementation from many shader languages.
// Used to find a given distance between two known locations (coordinates).
//
// The formula used here is taken from the Wikipedia article on the
// subject: https://en.wikipedia.org/wiki/Linear_interpolation#Programming_language_support
func Lerp(a, b, distance float32) float32 {
	return a + ((b - a) * distance)
}

// AI is an unbeatble computer controlled player.
func AI(paddle *Paddle, ball *Ball, elapsedTime float32) {
	if ball.Position.X <= (ball.Radius * 2) {
		return
	}

	var distance float32 = -(ball.Position.X - paddle.Position.X)
	var ghost = Ball{
		Position:   Position(ball.Position.X, ball.Position.Y),
		Dimensions: Dimension(0, 0, ball.Radius),
		Velocity:   Velocity(BallVelocity, BallVelocity),
	}

	for distance > 50 {
		ghost.Update(nil, nil, elapsedTime)

		distance = -(ghost.Position.X - paddle.Position.X)
	}

	if (ball.Position.Y < paddle.Position.Y) && (!paddle.IsTouchingTop() || paddle.IsTouchingBottom()) {
		if ghost.Position.Y < paddle.Position.Y {
			paddle.Position.Y -= paddle.Velocity.Y * elapsedTime
		}
	} else if (ball.Position.Y > paddle.Position.Y) && (!paddle.IsTouchingBottom() || paddle.IsTouchingTop()) {
		if ghost.Position.Y > paddle.Position.Y {
			paddle.Position.Y += paddle.Velocity.Y * elapsedTime
		}
	}
}

// Window is a convenience method for creating new windows using SDL2.
func Window(title string, width, height int32) *sdl.Window {
	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		width,
		height,
		sdl.WINDOW_SHOWN,
	)

	Abort(err)

	return window
}

// Renderer is a convenience method to create a 2D renderer with accelerated GPU
// support for the given window instance.
func Renderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	Abort(err)

	return renderer
}

// Texture is a convenience method to create an SDL pixel texture for the given
// renderer instance.
func Texture(renderer *sdl.Renderer, width, height int32) *sdl.Texture {
	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		width,
		height,
	)

	Abort(err)

	return texture
}

// Pixel will populate the given pixel in a set of pixels with the given colour.
func Pixel(x, y int32, colour entity.Colour, pixels []byte) {
	var index = ((y * WindowWidth) + x) * 4
	var bit = int32(len(pixels) - 4)

	if index < bit && index >= 0 {
		pixels[index] = colour.Red
		pixels[(index + 1)] = colour.Green
		pixels[(index + 2)] = colour.Blue
	}
}

// Clear will set all pixels in a given set of pixels to empty (black screen),
// iterating through in order that they are stored in memory.
func Clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

// Draw many objects on the screen in order of place in memory
func Draw(pixels []byte, objects [MaxObjects]Object) {
	Clear(pixels)

	for _, object := range objects {
		object.Draw(pixels)
	}
}

// Dimension will return a new dimension physics struct to define the size of a
// game object.
func Dimension(width, height, radius float32) entity.Dimensions {
	return entity.Dimensions{
		Width:  width,
		Height: height,
		Radius: radius,
	}
}

// Position provides a new position physics struct which depicts the current
// coordinates of the gane object this is assigned to.
func Position(x, y float32) entity.Position {
	return entity.Position{
		X: x,
		Y: y,
	}
}

// Velocity will return a physics struct which depicts the speed of a 2D game
// object in the world space.
func Velocity(x, y float32) entity.Velocity {
	return entity.Velocity{
		X: x,
		Y: y,
	}
}

// Centre will create a new position physics struct which points to the centre
// of the game screen window.
func Centre() entity.Position {
	return Position(float32((WindowWidth / 2)), float32((WindowHeight / 2)))
}

// White will return a Colour definition that can be used to populate a pixel as
// white using RGB.
func White() entity.Colour {
	return entity.Colour{
		Red:   255,
		Green: 255,
		Blue:  255,
	}
}

// IsKeyPressed checks if the given SDL keyboard scancode is actively being held
// or was pressed by the user.
func IsKeyPressed(scancode int) bool {
	return 0 != keyboard[scancode]
}

// WaitForPlayer will transition the game State from Waiting to Playing when the
// user hits the space bar key.
func WaitForPlayer() {
	if IsKeyPressed(sdl.SCANCODE_SPACE) && Waiting == State {
		State = Playing
	}
}

// Player returns a fully ready to use player object to be drawn and updated in
// the game world.
func Player(position entity.Position, colour entity.Colour) Paddle {
	return Paddle{
		Position:   position,
		Dimensions: Dimension(PaddleWidth, PaddleLength, 0),
		Velocity:   Velocity(0, PaddleVelocity),
		Colour:     colour,
		Score:      Tally(0, position, colour),
	}
}

// Pong will create a new ball instance that is ready to be drawn and updated in
// the game world.
func Pong(colour entity.Colour) Ball {
	return Ball{
		Position:   Centre(),
		Dimensions: Dimension(0, 0, BallRadius),
		Velocity:   Velocity(BallVelocity, BallVelocity),
		Colour:     colour,
	}
}

// Tally will create a new scoreboard at the given position's X cooridnate, on
// screen.
func Tally(score int, position entity.Position, colour entity.Colour) Score {
	position.X = Lerp(position.X, Centre().X, ScoreBoardOffsetX)
	position.Y = ScoreBoardOffsetY

	return Score{
		Number:   score,
		Position: position,
		Colour:   colour,
		Scale:    ScoreSize,
	}
}
