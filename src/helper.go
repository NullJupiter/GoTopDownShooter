package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type size struct {
	x, y int
}

type circle struct {
	x, y float64
	r    float64
}

type rectangle struct {
	x, y float64
	w, h float64
}

func textureFromBMP(renderer *sdl.Renderer, filepath string) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filepath)
	if err != nil {
		return nil, err
	}
	defer img.Free()
	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

/*func isColliding(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.x-c1.x, 2) + math.Pow(c2.y-c1.y, 2))
	return dist <= c1.radius+c2.radius
}*/

func isHover(mouseX, mouseY int32, buttonRect sdl.Rect) bool {
	if mouseX >= buttonRect.X && mouseX <= buttonRect.X+buttonRect.W {
		if mouseY >= buttonRect.Y && mouseY <= buttonRect.Y+buttonRect.H {
			return true
		}
	}
	return false
}
