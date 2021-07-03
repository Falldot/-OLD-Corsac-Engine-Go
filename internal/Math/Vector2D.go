package math

type Vector2D struct {
	X, Y float32
}

func (v1 *Vector2D) Add(v2 *Vector2D) {
	v1.X = +v2.X
	v1.Y = +v2.Y
}

func (v1 *Vector2D) Set(v2 *Vector2D) {
	v1.X = v2.X
	v1.Y = v2.Y
}
