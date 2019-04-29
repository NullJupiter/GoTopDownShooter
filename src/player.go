package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

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
	size              size
	collisionMask     circle
}

func newPlayer(x, y, movementSpeed float64, texture *sdl.Texture, textureRect sdl.Rect) (p playerStruct) {
	p.x, p.y = x, y
	p.movementSpeed = movementSpeed
	p.texture = texture
	p.textureRect = textureRect
	p.health = 100
	p.isAlive = true
	p.size.x = 32
	p.size.y = 32
	p.collisionMask = circle{x: x, y: y, r: float64(p.size.x / 2)}

	return p
}

func (p *playerStruct) update(dt float64, gameMap *mapStruct) {
	// look in direction of mouse
	dx := float64(mouseX) - p.x
	dy := float64(mouseY) - p.y
	p.angle = math.Atan2(dy, dx) * 180 / math.Pi

	// Shooting
	p.timeSinceLastShot += dt
	if mouseState == 1 {
		if p.timeSinceLastShot >= 0.3 {
			p.bullets = append(p.bullets, newBullet(p.texture, sdl.Rect{X: 32, Y: 0, W: 32, H: 32}, p.angle, p.x, p.y, 1000))
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
	// update collision mask
	p.collisionMask.x, p.collisionMask.y = p.x, p.y

	// update map collisions
	for _, tile := range gameMap.tiles {
		if tile.isCollider {
			collision, dir := p.collidesWithRectangle(tile.collsisionMask)
			if collision {
				switch dir {
				case 1:
					//p.x = tile.collsisionMask.x + tile.collsisionMask.w + p.collisionMask.r
					p.x += p.movementSpeed * dt
					break
				case 2:
					//p.x = tile.collsisionMask.x - p.collisionMask.r
					p.x -= p.movementSpeed * dt
					break
				case 3:
					//p.y = tile.collsisionMask.y + tile.collsisionMask.h + p.collisionMask.r
					p.y += p.movementSpeed * dt
					break
				case 4:
					//p.y = tile.collsisionMask.y - p.collisionMask.r
					p.y -= p.movementSpeed * dt
					break
				default:
					break
				}
			}
		}
	}
}

func (p *playerStruct) render(renderer *sdl.Renderer) {
	for i := 0; i < len(p.bullets); i++ {
		p.bullets[i].render(renderer)
	}

	renderer.CopyEx(
		p.texture,
		&p.textureRect,
		&sdl.Rect{X: int32(p.x - float64(p.size.x/2)), Y: int32(p.y - float64(p.size.y/2)), W: int32(p.size.x), H: int32(p.size.y)},
		p.angle+90,
		nil,
		sdl.FLIP_NONE)
}

func (p *playerStruct) collidesWithRectangle(rect rectangle) (bool, int) {
	if (p.collisionMask.r + rect.w/2) >= math.Sqrt(math.Pow(p.x-rect.x, 2)+math.Pow(p.y-rect.y, 2)) {
		playerBottom := p.collisionMask.y + p.collisionMask.r
		rectBottom := rect.y + rect.h
		playerRight := p.collisionMask.x + p.collisionMask.r
		rectRight := rect.x + rect.w

		bCollision := rectBottom - p.collisionMask.y - p.collisionMask.r
		tCollision := playerBottom - rect.y
		lCollision := playerRight - rect.x
		rCollision := rectRight - p.collisionMask.x - p.collisionMask.r

		if tCollision < bCollision && tCollision < lCollision && tCollision < rCollision {
			//Top collision
			return true, 4
		} else if bCollision < tCollision && bCollision < lCollision && bCollision < rCollision {
			//bottom collision
			return true, 3
		} else if lCollision < rCollision && lCollision < tCollision && lCollision < bCollision {
			//Left collision
			return true, 2
		} else if rCollision < lCollision && rCollision < tCollision && rCollision < bCollision {
			//Right collision
			return true, 1
		}
	}

	return false, 0
}
