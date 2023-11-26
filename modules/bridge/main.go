package main

// Мост позволяет разделить объект на две иерархии и подменять их

// создадим две иерархии:

// - Иерархия абстракции: сюда будут входить наши компьютеры
// - Иерархия реализации: сюда будут входить наши принтеры

// Эти две иерархии общаются между собой посредством Моста,
// в котором Абстракция (компьютер) содержит ссылку на Реализацию (принтер).
// И абстракцию, и реализацию можно разрабатывать отдельно, не влияя друг на друга.

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type AbstractComputer struct {
	printer Printer
}

func (c *AbstractComputer) SetPrinter(p Printer) {
	c.printer = p
}

func (c *AbstractComputer) Print() {
	fmt.Println("Print is not implemented")
}

type Mac struct {
	AbstractComputer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

type Windows struct {
	AbstractComputer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
