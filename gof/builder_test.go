package gof

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	normalBuilder, _ := getBuilder("normal")
	iglooBuilder, _ := getBuilder("igloo")

	normalBuilder.setNumFloor()
	normalBuilder.setWindowType()
	normalBuilder.setNumFloor()
	normalHouse := normalBuilder.getHouse()
	fmt.Printf("%+v\n", normalHouse)

	iglooBuilder.setNumFloor()
	iglooBuilder.setWindowType()
	iglooBuilder.setNumFloor()
	igloo := iglooBuilder.getHouse()
	fmt.Printf("%+v\n", igloo)
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) (IBuilder, error) {
	if builderType == "normal" {
		return newNormalBuilder(), nil
	}
	if builderType == "igloo" {
		return newIglooBuilder(), nil
	}
	return nil, fmt.Errorf("no builder of type %s found", builderType)
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}
