package cs

import (
	ecs "CRSC/internal/ECS"
	math "CRSC/internal/Math"
	"fmt"
	"io"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Map struct {
	texID                       string
	scale, tileSize, scaledSize int
	sizeY, sizeX                int
	path                        string
}

func (m *Map) Init() {
	var data []byte
	file, err := os.Open(m.path)
	if err != nil {
		panic(err)
	}
	n, err := file.Read(data)
	if err == io.EOF {
		panic(err)
	}
	fmt.Println(n)

}

func (m *Map) AddTile(src *sdl.Rect, pos *math.Vector2D) {
	e := ecs.CreateEntity()
	ecs.AddComponent(e,
		&Tile{src: src, position: pos},
		COMPONENT_TILE)
	ecs.GetComponent(e, COMPONENT_TILE).(*Tile).SetTex(m.texID)
}

func (m *Map) Update() {}
func (m *Map) Render() {}
