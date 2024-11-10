// Programmed by Piras314
// My first game with SDL in Go

// This program is part of the main package
package main

// Import the packages we will need for this project
import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Define constants, variables with constant values
const (
	screenWidth  = 500
	screenHeight = 600
)

// Function to easily make error messages so I don't have to repeat the same code so many times, somehow I haven't seen anyone else do this before
func errMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}

// Main Function of this program
func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}

	// Create an SDL Window Object
	window, err := sdl.CreateWindow(
		"Game in Go SDL2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL,
	)

	errMsg(err, "Initializing Window: ")

	// Destroy our window because we are using a c library so we don't have garbage cleanup
	defer window.Destroy()

	// Create a renderer with gpu accelerated graphics
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)

	errMsg(err, "Initializing Renderer: ")

	// Destroy our renderer
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	errMsg(err, "Creating Player: ")

	var enemies []basicEnemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy, err := newBasicEnemy(renderer, x, y)

			errMsg(err, "Creating basic enemy: ")
			enemies = append(enemies, enemy)
		}
	}

	// Main game loop
	for {
		// Check for events eg. mouseclick, mousemovement, keypress, keyrelease, window quit etc.
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			// If the x on the window is pressed (or equivalent in desktop env)
			case *sdl.QuitEvent:
				// Exit the program
				return
			}
		}
		// Set the draw colour to white
		renderer.SetDrawColor(255, 255, 255, 255)

		// Clear the screen with our draw colour
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		renderer.Present()

		// So we don't use too much cpu
		sdl.Delay(1)
	}
}
