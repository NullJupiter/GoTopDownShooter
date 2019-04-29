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
	mouseX     int32
	mouseY     int32
	mouseState uint32
)

func main() {
	// initialize sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("could not initialize go-sdl2: %v", err)
	}
	defer sdl.Quit()

	// create the window
	window, err := sdl.CreateWindow(
		"Top Down Shooter",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(fmt.Errorf("Could not create window. Error: %v", err))
	}
	defer window.Destroy()

	// put the game into fullscreen
	if err := window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP); err != nil {
		log.Fatalf("could not set window to fullscreen: %v", err)
	}

	// create the renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(fmt.Errorf("Could not create renderer. Error: %v", err))
	}
	defer renderer.Destroy()

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "linear")
	renderer.SetLogicalSize(windowWidth, windowHeight)

	// load the map
	gameMap, err := loadMap("assets/map.txt", renderer)
	if err != nil {
		log.Fatalf("could not load map from map file: %v", err)
	}

	// disable normal cursor
	if _, err = sdl.ShowCursor(0); err != nil {
		log.Fatalf("could not hide cursor: %v", err)
	}

	// load spriteSheet
	spriteSheet, err := textureFromBMP(renderer, "assets/spriteSheet.bmp")
	if err != nil {
		log.Fatalf("could not load spriteSheet: %v", err)
	}
	// load cursorTexture
	cursorTexture, err := textureFromBMP(renderer, "assets/cursor.bmp")
	if err != nil {
		log.Fatalf("could not load cursorTexture: %v", err)
	}
	// create cursor
	cursor := newCursor(cursorTexture, sdl.Rect{X: 0, Y: 0, W: 512, H: 512})
	// create player
	player := newPlayer(windowWidth/2-16, windowHeight/2-16, 350, spriteSheet, sdl.Rect{X: 0, Y: 0, W: 32, H: 32})

	// create variables for fps calculations
	var (
		newTime    time.Time
		dt         float64
		fpsCounter int
		fpsTime    float64
	)
	oldTime := time.Now()

	// enter the game loop
	running := true
	for running {
		// check for events
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
		mouseX, mouseY, mouseState = sdl.GetMouseState()

		// set background color
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// update and render the map tiles
		for _, tile := range gameMap.tiles {
			tile.update()
			tile.render(renderer)
		}

		// render the cursor
		cursor.render(renderer)

		// update and render the player
		player.update(dt, &gameMap)
		player.render(renderer)

		// present the scene
		renderer.Present()
	}
}
