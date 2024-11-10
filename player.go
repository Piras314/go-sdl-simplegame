// Programmed by Piras314
// My first game with SDL in Go

// This is part of the main package
package main

// Import the packages that we need for this program
import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Define constants
const (
	playerSpeed = 4
	playerSize  = 105
)

// Create the player struct
type player struct {
	tex  *sdl.Texture
	x, y float64
}

// Function to create a new player
func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	// Load the player sprite as a surface
	img, err := sdl.LoadBMP("res/img/player.bmp")

	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}

	defer img.Free()

	// Create a texture from that surface
	p.tex, err = renderer.CreateTextureFromSurface(img)

	if err != nil {
		return player{}, fmt.Errorf("creating player texture: %v", err)
	}

	// Set the player's position
	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize/2.0

	return p, nil
}

// Function to draw the player to the screen
func (p *player) draw(renderer *sdl.Renderer) {
	// Converting player coordinates to the top left of sprite
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0

	renderer.Copy(
		p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
	)
}

// Function to update the player
func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Move player left
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
}
