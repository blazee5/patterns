package main

import (
	"fmt"
)

type Car struct {
	brand       string
	model       string
	color       string
	engine      string
	doors       int
	maxSpeed    int
	isAutomatic bool
}

func (c *Car) String() string {
	transmission := "Manual"
	if c.isAutomatic {
		transmission = "Automatic"
	}
	return fmt.Sprintf("%s %s (%s) - %s engine, %d doors, %d km/h, %s",
		c.brand, c.model, c.color, c.engine, c.doors, c.maxSpeed, transmission)
}

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: &Car{
			doors:    4,
			maxSpeed: 180,
		},
	}
}

func (cb *CarBuilder) WithBrand(brand string) *CarBuilder {
	cb.car.brand = brand
	return cb
}

func (cb *CarBuilder) WithModel(model string) *CarBuilder {
	cb.car.model = model
	return cb
}

func (cb *CarBuilder) WithColor(color string) *CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *CarBuilder) WithEngine(engine string) *CarBuilder {
	cb.car.engine = engine
	return cb
}

func (cb *CarBuilder) WithDoors(doors int) *CarBuilder {
	cb.car.doors = doors
	return cb
}

func (cb *CarBuilder) WithMaxSpeed(maxSpeed int) *CarBuilder {
	cb.car.maxSpeed = maxSpeed
	return cb
}

func (cb *CarBuilder) WithTransmission(isAutomatic bool) *CarBuilder {
	cb.car.isAutomatic = isAutomatic
	return cb
}

func (cb *CarBuilder) Build() *Car {
	return cb.car
}

func main() {
	sportsCar := NewCarBuilder().
		WithBrand("Ferrari").
		WithModel("488 GTB").
		WithColor("Red").
		WithEngine("V8").
		WithDoors(2).
		WithMaxSpeed(330).
		WithTransmission(true).
		Build()

	electricCar := NewCarBuilder().
		WithBrand("Tesla").
		WithModel("Model S").
		WithColor("Blue").
		WithEngine("Electric").
		WithDoors(4).
		WithMaxSpeed(250).
		WithTransmission(true).
		Build()

	fmt.Println("Sports Car:", sportsCar)
	fmt.Println("Electric Car:", electricCar)
}
