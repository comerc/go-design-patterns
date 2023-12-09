// Посетитель позволяет добавить новую операцию для целой иерархии классов,
// не изменяя код этих классов (нельзя заменить простой перегрузкой методов)
// https://refactoring.guru/ru/design-patterns/visitor-double-dispatch

// реализация паттерна Visitor через дженерик!

package main

import "fmt"

type Visitor[T Element] interface {
	Visit(element *T)
}

type Acceptor interface {
	Accept(visitor Visitor[Element])
}

type Element struct {
	name string
}

func (ce *Element) Accept(visitor Visitor[Element]) {
	visitor.Visit(ce)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) Visit(element *Element) {
	fmt.Println("Calculating area for", element.name)
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) Visit(element *Element) {
	fmt.Println("Calculating coordinates for", element.name)
}

func main() {
	square := Element{name: "Square"}
	circle := Element{name: "Circle"}
	rectangle := Element{name: "Rectangle"}

	areaCalculator := &AreaCalculator{}
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
