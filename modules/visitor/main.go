package main

import "fmt"

// Посетитель позволяет добавить новую операцию для целой иерархии классов,
// не изменяя код этих классов (нельзя заменить простой перегрузкой методов)
// https://refactoring.guru/ru/design-patterns/visitor-double-dispatch

// реализация через дженерик паттерна Visitor
// смущает IsShape - предложите решение лучше?

type Shape interface {
	IsShape()
}

type Visitor[T Shape] interface {
	visit(s T)
}

type Square struct {
	side int
}

func (*Square) IsShape() {}

type Circle struct {
	radius int
}

func (*Circle) IsShape() {}

type Rectangle struct {
	l int
	b int
}

func (*Rectangle) IsShape() {}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visit(s Shape) {
	switch s.(type) {
	case *Square:
		fmt.Println("Calculating area for square")
	case *Circle:
		fmt.Println("Calculating area for circle")
	case *Rectangle:
		fmt.Println("Calculating area for rectangle")
	}
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visit(s Shape) {
	switch s.(type) {
	case *Square:
		fmt.Println("Calculating middle point coordinates for square")
	case *Circle:
		fmt.Println("Calculating middle point coordinates for circle")
	case *Rectangle:
		fmt.Println("Calculating middle point coordinates for rectangle")
	}
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}
	areaCalculator.visit(square)
	areaCalculator.visit(circle)
	areaCalculator.visit(rectangle)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	middleCoordinates.visit(square)
	middleCoordinates.visit(circle)
	middleCoordinates.visit(rectangle)
}
