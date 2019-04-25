package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cursorWidth  = 512 / 8
	cursorHeight = 512 / 8
)

type cursorStruct struct {
	texture     *sdl.Texture
	textureRect sdl.Rect
}

func newCursor(texture *sdl.Texture, textureRect sdl.Rect) (c cursorStruct) {
	c.texture = texture
	c.textureRect = textureRect

	return c
}

func (c *cursorStruct) render(renderer *sdl.Renderer) {
	renderer.Copy(
		c.texture,
		&c.textureRect,
		&sdl.Rect{X: mouseX, Y: mouseY, W: cursorWidth, H: cursorHeight})
}
