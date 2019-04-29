package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type tileStruct struct {
	x, y           float64
	texture        *sdl.Texture
	textureRect    sdl.Rect
	isCollider     bool
	collsisionMask rectangle
}

func (t *tileStruct) update() {

}

func (t *tileStruct) render(renderer *sdl.Renderer) {
	renderer.Copy(t.texture, &t.textureRect, &sdl.Rect{X: int32(t.x - t.collsisionMask.w/2), Y: int32(t.y - t.collsisionMask.h/2), W: 32, H: 32})
}

type mapStruct struct {
	tiles []tileStruct
}

const tileSize = 32

func loadMap(filepath string, renderer *sdl.Renderer) (mapStruct, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return mapStruct{}, err
	}
	defer f.Close()

	tileSheet, err := textureFromBMP(renderer, "assets/tileSheet.bmp")
	if err != nil {
		return mapStruct{}, fmt.Errorf("could not load tileSheet: %v", err)
	}

	var tiles []tileStruct
	var i int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for j, char := range strings.Split(line, "") {
			var tile tileStruct
			var tileRect sdl.Rect
			if char == "1" {
				tileRect = sdl.Rect{X: 0, Y: 0, W: tileSize, H: tileSize}
				tile = tileStruct{x: float64(j*tileSize + tileSize/2), y: float64(i*tileSize + tileSize/2), texture: tileSheet, textureRect: tileRect, isCollider: true}
			} else if char == "0" {
				tileRect = sdl.Rect{X: 32, Y: 0, W: tileSize, H: tileSize}
				tile = tileStruct{x: float64(j*tileSize + tileSize/2), y: float64(i*tileSize + tileSize/2), texture: tileSheet, textureRect: tileRect, isCollider: false}
			} else {
				return mapStruct{}, fmt.Errorf("invalid tile option in map file: %v", err)
			}
			tile.collsisionMask = rectangle{x: tile.x, y: tile.y, w: tileSize, h: tileSize}
			tiles = append(tiles, tile)
		}
		i++
	}

	return mapStruct{tiles: tiles}, nil
}
