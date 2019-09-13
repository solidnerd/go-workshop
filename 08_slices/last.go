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

	computerStore.Computers = append(computerStore.Computers, &Computer{
		Brand: "Apple",
		Model: "Macbook Pro",
		Price: 4000,
	})

	computerStore.Computers = append(computerStore.Computers, &Computer{
		Brand: "Apple ",
		Model: "iPhone 11 Max",
		Price: 2000,
	})

	for count, computer := range computerStore.Computers {
		fmt.Printf("Computer %d \n", count+1)
		computer.Describe()
	}

}
