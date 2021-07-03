/////////////////////////////////////////////////////
//
// 	ecs @ 2021
//	fe.offing@gmail.com
//
/////////////////////////////////////////////////////
package cs

import (
	ecs "CRSC/internal/ECS"
	math "CRSC/internal/Math"
)

type Transform struct {
	Entity   *ecs.Entity
	rotate   *math.Vector2D
	Position *math.Vector2D
	Velocity *math.Vector2D

	Height, Width, Scale, speed int
}

func CreateTransform(x float32, y float32, w int, h int) *Transform {
	return &Transform{Position: &math.Vector2D{X: x, Y: y}, Width: w, Height: h}
}

func (t *Transform) Init(e *ecs.Entity) {
	t.Entity = e
	t.Velocity = &math.Vector2D{X: 0, Y: 0}
	t.Scale = 2
}

func (t *Transform) Translate(v *math.Vector2D) {
	t.Position.Add(v)
}

func (t *Transform) Rotate(v *math.Vector2D) {
	t.rotate.Add(v)
}

func (t *Transform) SetScale(value int) {
	t.Scale = value
}

func (t *Transform) Speed(value int) {
	t.speed = value
}

func (t *Transform) Update() {
	t.Position.X += t.Velocity.X * float32(t.speed)
	t.Position.Y += t.Velocity.Y * float32(t.speed)
}
func (t *Transform) Render() {}
