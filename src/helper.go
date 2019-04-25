package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func textureFromBMP(renderer *sdl.Renderer, filepath string) *sdl.Texture {
	img, err := sdl.LoadBMP(filepath)
	if err != nil {
		panic(fmt.Errorf("Could not load bmp file (%v). Error: %v", filepath, err))
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Could not create texture from surface (%v). Error: %v", filepath, err))
	}

	return texture
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

type circle struct {
	x, y   float64
	radius float64
}

func isColliding(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.x-c1.x, 2) + math.Pow(c2.y-c1.y, 2))
	return dist <= c1.radius+c2.radius
}

func isHover(mouseX, mouseY int32, buttonRect sdl.Rect) bool {
	if mouseX >= buttonRect.X && mouseX <= buttonRect.X+buttonRect.W {
		if mouseY >= buttonRect.Y && mouseY <= buttonRect.Y+buttonRect.H {
			return true
		}
	}
	return false
}
