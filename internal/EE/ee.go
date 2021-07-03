package ee

import "github.com/veandco/go-sdl2/sdl"

var State []uint8

func Init() {
	State = sdl.GetKeyboardState()
}
