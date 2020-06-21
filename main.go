package main

import (
	"fmt"
	"time"

	"github.com/jordanbrauer/gogame/game"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var err error = sdl.Init(sdl.INIT_EVERYTHING)

	game.Abort(err)

	var window *sdl.Window = game.Window("Pong", game.WindowWidth, game.WindowHeight)
	var renderer *sdl.Renderer = game.Renderer(window)
	var texture *sdl.Texture = game.Texture(renderer, game.WindowWidth, game.WindowHeight)

	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	defer texture.Destroy()

	var pixels = make([]byte, (game.WindowWidth * game.WindowHeight * 4))
	var uiColour = game.White()
	var player = game.Player(game.Position(50, (float32(game.WindowHeight)-100)), uiColour)
	var computer = game.Player(game.Position((float32(game.WindowWidth)-50), 100), uiColour)
	var ball = game.Pong(uiColour)
	var objects = [game.MaxObjects]game.Object{
		&player,
		&player.Score,
		&computer,
		&computer.Score,
		&ball,
	}
	var frameStart time.Time
	var frameElapsed float32
	var running = true

	fmt.Println("Welcome to Pong!")

	for running {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false

				fmt.Println("See ya' later.")

				break
			}
		}

		switch game.State {
		case game.Waiting:
			game.WaitForPlayer()

			break
		case game.Playing:
			player.Update(frameElapsed)
			ball.Update(&player, &computer, frameElapsed)
			game.AI(&computer, &ball, frameElapsed)

			break
		}

		game.Draw(pixels, objects)
		texture.Update(nil, pixels, (int(game.WindowWidth) * 4))
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		frameElapsed = float32(time.Since(frameStart).Seconds())

		if frameElapsed < 0.005 {
			sdl.Delay(5 - uint32((frameElapsed * 1000.0)))
			frameElapsed = float32(time.Since(frameStart).Seconds())
		}
	}
}
