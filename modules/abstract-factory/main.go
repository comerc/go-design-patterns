package main

import "fmt"

// позволяет создавать группы связанных между собой объектов,
// при этом не привязываться к классам создаваемых объетов

// преимущества:
// - сочетаемость создаваемых объектов (могут быть горизонтальные связи)
// - избавляет клиентский код от привязки к конкретным объектам
// - выдеяет код в одном месте (т.е. упрощает поддержку)
// - легче добавлять новые типы объектов
// - реализует принцип открытости-закрытости из SOLID
// недостатки:
// - усложняет код (требуется реализовывать много объектов)
// - требует заранее знать единый интерфейс возвращаемого типа

type ISportsFactory interface {
	makeShoe(size int) IShoe
	makeShirt(size int) IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

type Adidas struct{}

func (a *Adidas) makeShoe(size int) IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: size,
		},
	}
}

func (a *Adidas) makeShirt(size int) IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: size,
		},
	}
}

type Nike struct{}

func (n *Nike) makeShoe(size int) IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: size,
		},
	}
}

func (n *Nike) makeShirt(size int) IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: size,
		},
	}
}

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	Logo() string
	Size() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) Logo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) Size() int {
	return s.size
}

type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	Logo() string
	Size() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) Logo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) Size() int {
	return s.size
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe(10)
	nikeShirt := nikeFactory.makeShirt(11)

	adidasShoe := adidasFactory.makeShoe(10)
	adidasShirt := adidasFactory.makeShirt(11)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.Logo())
	fmt.Println()
	fmt.Printf("Size: %d", s.Size())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.Logo())
	fmt.Println()
	fmt.Printf("Size: %d", s.Size())
	fmt.Println()
}
