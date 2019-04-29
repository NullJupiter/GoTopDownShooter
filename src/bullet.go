package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletWidth  = 32
	bulletHeight = 16
)

type bulletStruct struct {
	x, y          float64
	movementSpeed float64
	isAlive       bool
	angle         float64
	texture       *sdl.Texture
	textureRect   sdl.Rect
	xDir, yDir    float64
}

func newBullet(texture *sdl.Texture, textureRect sdl.Rect, angle, x, y, movementSpeed float64) (b bulletStruct) {
	b.angle = angle
	b.x, b.y = x, y
	b.movementSpeed = movementSpeed
	b.texture = texture
	b.textureRect = textureRect
	b.xDir = float64(mouseX) - b.x
	b.yDir = float64(mouseY) - b.y
	magnitude := math.Sqrt(math.Pow(b.xDir, 2) + math.Pow(b.yDir, 2))
	b.xDir /= magnitude
	b.yDir /= magnitude

	return b
}

func (b *bulletStruct) update(dt float64) {
	b.x += b.movementSpeed * dt * b.xDir
	b.y += b.movementSpeed * dt * b.yDir
}

func (b *bulletStruct) render(renderer *sdl.Renderer) {
	renderer.CopyEx(
		b.texture,
		&b.textureRect,
		&sdl.Rect{X: int32(b.x - bulletWidth/2), Y: int32(b.y - bulletHeight/2), W: bulletWidth, H: bulletHeight},
		b.angle,
		nil,
		sdl.FLIP_NONE)
}
