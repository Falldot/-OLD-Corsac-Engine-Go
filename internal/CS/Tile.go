package cs

import (
	ecs "CRSC/internal/ECS"
	ge "CRSC/internal/GE"
	math "CRSC/internal/Math"

	"github.com/veandco/go-sdl2/sdl"
)

type Tile struct {
	tex      *sdl.Texture
	src, dst *sdl.Rect
	position *math.Vector2D
}

func (t *Tile) SetTex(name string) {
	t.tex = ge.Load(name)
}

func (t *Tile) Init(e *ecs.Entity) {}

func (t *Tile) Update() {
	t.dst.X = int32(t.position.X)
	t.dst.Y = int32(t.position.Y)
}

func (t *Tile) Render() {
	err := ge.Render.Copy(t.tex, t.src, t.dst)
	if err != nil {
		panic(err)
	}
}
