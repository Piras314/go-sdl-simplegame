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
	basicEnemySize = 105
)

// Create the basicEnemy struct
type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

// Function to create a new basicEnemy
func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error) {
	// Load the basic enemy sprite as a surface
	img, err := sdl.LoadBMP("res/img/basic_enemy.bmp")

	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading basic enemy sprite: %v", err)
	}

	defer img.Free()

	// Create a texture from that surface
	be.tex, err = renderer.CreateTextureFromSurface(img)

	if err != nil {
		return basicEnemy{}, fmt.Errorf("creating basic enemy texture: %v", err)
	}

	// Set the basic enemy's position
	be.x = x
	be.y = y

	return be, nil
}

// Function do draw the basic enemy to the screen
func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	// Converting player coordinates to the top left of sprite
	x := be.x - basicEnemySize/2.0
	y := be.y - basicEnemySize/2.0

	renderer.CopyEx(
		be.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
		180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE,
	)
}
