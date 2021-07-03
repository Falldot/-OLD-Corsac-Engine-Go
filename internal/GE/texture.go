package ge

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func Load(path string) *sdl.Texture {
	s, err := img.Load(path)
	if err != nil {
		panic(err)
	}
	t, err := Render.CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}
	s.Free()
	return t
}
