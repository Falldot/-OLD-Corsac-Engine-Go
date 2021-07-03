package cs

import (
	ecs "CRSC/internal/ECS"
	ge "CRSC/internal/GE"
	math "CRSC/internal/Math"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	Transform interface{}

	tex      *sdl.Texture
	src, dst *sdl.Rect

	animated bool
	frame    int
	speed    int

	animIndex  int
	animations map[string]*math.Animation

	flip sdl.RendererFlip
}

func (s *Sprite) SetTex(name string) *Sprite {
	s.tex = ge.Load(name)
	return s
}

func (s *Sprite) Init(e *ecs.Entity) {
	s.Transform = ecs.GetComponent(e, COMPONENT_TRANSFORM)
	s.speed = 100
	s.src = &sdl.Rect{}
	s.dst = &sdl.Rect{}
	s.src.W = int32(s.Transform.(*Transform).Width)
	s.src.H = int32(s.Transform.(*Transform).Height)
	s.animations = make(map[string]*math.Animation)
}

func (s *Sprite) Update() {
	if s.animated {
		s.src.X = s.src.W * (int32(sdl.GetTicks()) / int32(s.speed) % int32(s.frame))
	}

	s.src.Y = int32(s.animIndex) * int32(s.Transform.(*Transform).Height)

	s.dst.X = int32(s.Transform.(*Transform).Position.X)
	s.dst.Y = int32(s.Transform.(*Transform).Position.Y)
	s.dst.W = int32(s.Transform.(*Transform).Width) * int32(s.Transform.(*Transform).Scale)
	s.dst.H = int32(s.Transform.(*Transform).Height) * int32(s.Transform.(*Transform).Scale)
}

func (s *Sprite) Play(name string) {
	s.frame = s.animations[name].Frames
	s.animIndex = s.animations[name].Index
	s.speed = s.animations[name].Speed
}

func (s *Sprite) AddAnimation(name string, x int, y int, speed int) *Sprite {
	s.animated = true
	anim := &math.Animation{Index: x, Frames: y, Speed: speed}
	s.animations[name] = anim
	if name == "Idle" {
		s.Play("Idle")
	}
	return s
}

func (s *Sprite) Render() {
	err := ge.Render.CopyEx(s.tex, s.src, s.dst, 0, nil, s.flip)
	if err != nil {
		panic(err)
	}
}
