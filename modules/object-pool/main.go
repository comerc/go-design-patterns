package main

// В этом примере, `Pool` - это структура, которая управляет двумя списками объектов:
// `available` (доступные) и `inUse` (используемые).
// Когда клиент запрашивает объект из пула с помощью метода `Borrow`,
// объект перемещается из списка `available` в список `inUse`.
// Когда объект возвращается в пул с помощью метода `Return`,
// он перемещается обратно в список `available`.

import (
	"fmt"
	"sync"
)

type Object struct {
	id string
}

type Pool struct {
	available []*Object
	inUse     []*Object
	mtx       sync.Mutex
}

func NewPool(objects []*Object) *Pool {
	return &Pool{
		available: objects,
		inUse:     make([]*Object, 0),
	}
}

func (p *Pool) Borrow() *Object {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if len(p.available) == 0 {
		return nil
	}

	obj := p.available[0]
	p.available = p.available[1:]
	p.inUse = append(p.inUse, obj)

	return obj
}

func (p *Pool) Return(obj *Object) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	for i, o := range p.inUse {
		if o == obj {
			p.inUse = append(p.inUse[:i], p.inUse[i+1:]...)
			p.available = append(p.available, obj)
			break
		}
	}
}

func main() {
	objects := []*Object{{id: "1"}, {id: "2"}, {id: "3"}}
	pool := NewPool(objects)

	obj := pool.Borrow()
	fmt.Println("Borrowed object", obj.id)

	pool.Return(obj)
	fmt.Println("Returned object", obj.id)
}
