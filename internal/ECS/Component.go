/////////////////////////////////////////////////////
//
// 	ecs @ 2021
//	fe.offing@gmail.com
//
/////////////////////////////////////////////////////

package ecs

import (
	"sort"
)

type data interface {
	Init(Entity *Entity)
	Update()
	Render()
}

type component struct {
	Entity        *Entity
	componentType int
	data          data
}

///////////////////////
// Методы игрового цикла
func (c *component) Init(Entity *Entity) {
	c.data.Init(Entity)
}

func (c *component) Update() {
	c.data.Update()
}

func (c *component) Render() {
	c.data.Render()
}

///////////////////////

// AddComponent - добавить компонент сущности
func AddComponent(e *Entity, d data, tp int) {
	e.components = append(e.components, &component{e, tp, d})
}

// GetComponent - получить компонент сущности
func GetComponent(e *Entity, tp int) data {
	return e.components[sort.Search(len(e.components), func(i int) bool { return tp <= e.components[i].componentType })].data
}

// DelComponent - удалить компонент сущности
func DelComponent(e *Entity, tp int) {
	i := sort.Search(len(e.components), func(i int) bool { return tp <= e.components[i].componentType })
	e.components[i] = e.components[len(e.components)-1]
	e.components = e.components[:len(e.components)-1]
}
