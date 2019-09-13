package main

import "fmt"

type Computer struct {
	Brand string
	Model string
	Price int
}

type ComputerStore struct {
	Computers []*Computer
}

func (c *Computer) Describe() {
	fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

func main() {

	computerStore := &ComputerStore{}
	computerStore.Computers = []*Computer{}

	computerStore.Computers = append(computerStore.Computers, &Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	})

	//TODO: Create an Loop to iterate through the slice and describe every added computer
}
