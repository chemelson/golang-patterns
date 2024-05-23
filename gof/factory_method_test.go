package gof

import (
	"fmt"
	"testing"
)

// It is impossible to implement canonical Factory Method in Go, because it lacks OOP features.
// But we will implement Simple Factory

func TestFactoryMethod(t *testing.T) {
	ak47, _ := getGun("ak47")
	fmt.Printf("%T : %+v\n", ak47, ak47)

	fmt.Println(ak47.getName())
	fmt.Println(ak47.getPower())

	fmt.Println("---------")

	musket, _ := getGun("musket")
	fmt.Printf("%T : %+v\n", musket, musket)

	fmt.Println(musket.getName())
	fmt.Println(musket.getPower())
}

// Abstract product
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Concrete products

// This is part of composition
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK 47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// And here is factory itself
func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("unknown type")
	}
}
