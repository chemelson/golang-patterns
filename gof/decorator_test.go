package gof

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &PlainPizza{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price: %d\n", pizzaWithCheeseAndTomato.getPrice())
}

// Abstract component
type IPizza interface {
	getPrice() int
}

// Concrete component
type PlainPizza struct {
}

func (p *PlainPizza) getPrice() int {
	return 15
}

// First decorator
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// Second decorator
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
