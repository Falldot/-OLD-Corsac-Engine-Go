package ge

import (
	math "CRSC/internal/Math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var windowSize *math.Vector2D
var Window *sdl.Window
var Render *sdl.Renderer

func Init(name string, x float32, y float32) {
	windowSize = &math.Vector2D{X: x, Y: y}
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err = img.Init(img.INIT_PNG); err != nil {
		panic(err)
	}
	Window, err = sdl.CreateWindow(name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(windowSize.X), int32(windowSize.Y), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	Render, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
}

func Quit() {
	Render.Destroy()
	Window.Destroy()
	sdl.Quit()
}
