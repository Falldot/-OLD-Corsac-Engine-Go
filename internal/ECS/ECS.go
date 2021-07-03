/////////////////////////////////////////////////////
//
// 	ecs @ 2021
//	fe.offing@gmail.com
//
/////////////////////////////////////////////////////
package ecs

import (
	"sort"
	"sync/atomic"
)

var idIncEntity uint64

type Entity struct {
	id         uint64
	parent     *Entity
	childrens  []*Entity
	components []*component

	active bool
}

var EntityCollection []*Entity

// Active - активировать или диактивировать сущность
// Диактивированная сущность не обновляеться и ресуеться
func (e *Entity) Active(v bool) {
	e.active = v
}

// CreateEntity - создать сущность
func CreateEntity() *Entity {
	newElem := Entity{id: atomic.AddUint64(&idIncEntity, 1)}
	EntityCollection = append(EntityCollection, &newElem)
	return &newElem
}

// RemoveEntity - удалить сущность
func RemoveEntity(e *Entity) {
	i := sort.Search(len(EntityCollection), func(i int) bool { return e.id <= EntityCollection[i].id })
	EntityCollection = append(EntityCollection[:i], EntityCollection[i+1:]...)
}

// AddChildren - привязать к сущности дочернию сущность
func AddChildren(e *Entity, es *Entity) {
	es.parent = es
	e.childrens = append(e.childrens, es)
}

// SetParent - установить родительскую сущность
func SetParent(e1 *Entity, e2 *Entity) {
	e1.parent = e2
}

/////////////////////////////////////////////////////
// Методы игрового цикла
func Init() {
	for _, v := range EntityCollection {
		for _, c := range v.components {
			c.Init(v)
		}
	}
}
func Update() {
	for _, v := range EntityCollection {
		for _, c := range v.components {
			if v.active {
				c.Update()
			}
		}
	}
}
func Render() {
	for _, v := range EntityCollection {
		for _, c := range v.components {
			if v.active {
				c.Render()
			}
		}
	}
}

/////////////////////////////////////////////////////
