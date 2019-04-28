package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSize = 50

type playerStruct struct {
	x, y              float64
	isAlive           bool
	textureRect       sdl.Rect
	texture           *sdl.Texture
	movementSpeed     float64
	health            int
	angle             float64
	bullets           []bulletStruct
	timeSinceLastShot float64
}

func newPlayer(x, y, movementSpeed float64, texture *sdl.Texture, textureRect sdl.Rect) (p playerStruct) {
	p.x, p.y = x, y
	p.movementSpeed = movementSpeed
	p.texture = texture
	p.textureRect = textureRect
	p.health = 100
	p.isAlive = true

	return p
}

func (p *playerStruct) update(dt float64) {
	// look in direction of mouse
	dx := float64(mouseX) - p.x
	dy := float64(mouseY) - p.y
	p.angle = math.Atan2(dy, dx) * 180 / math.Pi

	// Shooting
	p.timeSinceLastShot += dt
	if mouseState == 1 {
		if p.timeSinceLastShot >= 0.3 {
			p.bullets = append(p.bullets, newBullet(p.texture, sdl.Rect{X: 32, Y: 0, W: 32, H: 32}, p.angle, p.x, p.y, 10))
			p.timeSinceLastShot = 0
		}
	}

	// update bullets
	for i := 0; i < len(p.bullets); i++ {
		p.bullets[i].update(dt)
	}

	// WASD Movement
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_W] == 1 {
		p.y -= p.movementSpeed * dt
	} else if keys[sdl.SCANCODE_A] == 1 {
		p.x -= p.movementSpeed * dt
	} else if keys[sdl.SCANCODE_S] == 1 {
		p.y += p.movementSpeed * dt
	} else if keys[sdl.SCANCODE_D] == 1 {
		p.x += p.movementSpeed * dt
	}
}

func (p *playerStruct) render(renderer *sdl.Renderer) {
	for i := 0; i < len(p.bullets); i++ {
		p.bullets[i].render(renderer)
	}

	renderer.CopyEx(
		p.texture,
		&p.textureRect,
		&sdl.Rect{X: int32(p.x - playerSize/2), Y: int32(p.y - playerSize/2), W: playerSize, H: playerSize},
		p.angle+90,
		nil,
		sdl.FLIP_NONE)
}
