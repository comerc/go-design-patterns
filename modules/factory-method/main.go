package main

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	Name() string
	Power() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) Name() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) Power() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() *Ak47 {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() *Musket {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// фабричный метод (или виртуальный конструктор)
// плюсы:
// - IGun избавляет от привязки к конкретному типу,
// - общий конструктор упрощает добавление новых типов,
// - реализует принцип открытости-закрытости из SOLID
// минусы:
// - раздуваемый IGun (много методов)
// - "божественный" getGun()
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.Name())
	fmt.Println()
	fmt.Printf("Power: %d", g.Power())
	fmt.Println()
}
