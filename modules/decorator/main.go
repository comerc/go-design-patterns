package main

// Декоратор или обёртка (wrapper)

import "fmt"

type Pizza interface {
	Price() int
}

type VeggieMania struct{}

func (p *VeggieMania) Price() int {
	return 15
}

type TomatoTopping struct {
	pizza Pizza
}

func (c *TomatoTopping) Price() int {
	pizzaPrice := c.pizza.Price()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) Price() int {
	pizzaPrice := c.pizza.Price()
	return pizzaPrice + 10
}

func main() {
	pizza := &VeggieMania{}
	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.Price())

	// или
	// pizza := &TomatoTopping{&CheeseTopping{&VeggieMania{}}}
	// fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizza.Price())
}
