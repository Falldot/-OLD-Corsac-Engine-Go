package cs

import (
	ecs "CRSC/internal/ECS"
	ee "CRSC/internal/EE"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardControlle struct {
	Transform interface{}
	Sprite    interface{}

	W, A, S, D bool
}

func (k *KeyboardControlle) Init(e *ecs.Entity) {
	k.Transform = ecs.GetComponent(e, COMPONENT_TRANSFORM)
	k.Sprite = ecs.GetComponent(e, COMPONENT_SPRITE)
}

func (k *KeyboardControlle) Update() {
	if ee.State[sdl.SCANCODE_W] != 0 {
		k.W = true
	} else {
		k.W = false
	}
	if ee.State[sdl.SCANCODE_S] != 0 {
		k.S = true
	} else {
		k.S = false
	}
	if ee.State[sdl.SCANCODE_A] != 0 {
		k.A = true
	} else {
		k.A = false
	}
	if ee.State[sdl.SCANCODE_D] != 0 {
		k.D = true
	} else {

		k.D = false
	}
	if !k.W && !k.S && !k.A && !k.D {
		k.Sprite.(*Sprite).Play("Idle")
		k.Transform.(*Transform).Velocity.X = 0
		k.Transform.(*Transform).Velocity.Y = 0
	} else if k.D && k.S {
		k.Sprite.(*Sprite).Play("DownRight")
		k.Sprite.(*Sprite).flip = sdl.FLIP_NONE
		k.Transform.(*Transform).Velocity.X = 1
		k.Transform.(*Transform).Velocity.Y = 1
	} else if k.A && k.S {
		k.Sprite.(*Sprite).Play("DownRight")
		k.Sprite.(*Sprite).flip = sdl.FLIP_HORIZONTAL
		k.Transform.(*Transform).Velocity.X = -1
		k.Transform.(*Transform).Velocity.Y = 1
	} else if k.W && k.A {
		k.Sprite.(*Sprite).Play("Right")
		k.Sprite.(*Sprite).flip = sdl.FLIP_HORIZONTAL
		k.Transform.(*Transform).Velocity.X = -1
		k.Transform.(*Transform).Velocity.Y = -1
	} else if k.W && k.D {
		k.Sprite.(*Sprite).Play("Right")
		k.Sprite.(*Sprite).flip = sdl.FLIP_NONE
		k.Transform.(*Transform).Velocity.X = 1
		k.Transform.(*Transform).Velocity.Y = -1
	} else if k.W {
		k.Sprite.(*Sprite).Play("Up")
		k.Transform.(*Transform).Velocity.X = -1
		k.Transform.(*Transform).Velocity.Y = 0
	} else if k.S {
		k.Sprite.(*Sprite).Play("Down")
		k.Transform.(*Transform).Velocity.X = 1
		k.Transform.(*Transform).Velocity.Y = 0
	} else if k.A {
		k.Sprite.(*Sprite).Play("Right")
		k.Sprite.(*Sprite).flip = sdl.FLIP_HORIZONTAL
		k.Transform.(*Transform).Velocity.X = 1
		k.Transform.(*Transform).Velocity.Y = 0
	} else if k.D {
		k.Sprite.(*Sprite).Play("Right")
		k.Sprite.(*Sprite).flip = sdl.FLIP_NONE
		k.Transform.(*Transform).Velocity.X = 1
		k.Transform.(*Transform).Velocity.Y = 0
	}
}

func (k *KeyboardControlle) Render() {}
