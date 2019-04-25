package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 1920
	windowHeight = 1080
)

var (
	mouseX int32
	mouseY int32
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("could not initialize go-sdl2: %v", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Top Down Shooter",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_OPENGL /*|sdl.WINDOW_ALLOW_HIGHDPI*/)
	if err != nil {
		panic(fmt.Errorf("Could not create window. Error: %v", err))
	}
	defer window.Destroy()

	if err := window.SetFullscreen(sdl.WINDOW_FULLSCREEN); err != nil {
		log.Fatalf("could not set window to fullscreen: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(fmt.Errorf("Could not create renderer. Error: %v", err))
	}
	defer renderer.Destroy()

	if _, err = sdl.ShowCursor(0); err != nil {
		log.Fatalf("could not hide cursor: %v", err)
	}

	spriteSheet := textureFromBMP(renderer, "assets/player.bmp")
	cursorTexture := textureFromBMP(renderer, "assets/cursor.bmp")
	cursor := newCursor(cursorTexture, sdl.Rect{X: 0, Y: 0, W: 512, H: 512})
	player := newPlayer(0, 0, 350, spriteSheet, sdl.Rect{X: 0, Y: 0, W: 32, H: 32})

	var (
		newTime    time.Time
		dt         float64
		fpsCounter int
		fpsTime    float64
	)
	oldTime := time.Now()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		// update frame counter and delta time
		newTime = time.Now()
		dt = newTime.Sub(oldTime).Seconds()
		oldTime = newTime
		// print current fps
		fpsTime += dt
		if fpsTime >= 1 {
			fmt.Println("\033cFPS:", fpsCounter)
			fpsTime = 0
			fpsCounter = 0
		} else {
			fpsCounter++
		}

		// update mouse position
		mouseX, mouseY, _ = sdl.GetMouseState()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		cursor.render(renderer)

		player.update(dt)
		player.render(renderer)

		renderer.Present()
	}
}
