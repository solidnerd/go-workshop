package main

import "fmt"

type Computer struct {
	Brand string
	Model string
	Price int
}

// TODO: Create a "Ask" method that answers or throw an an error if the question is longer than 20 characters and handle your two return values
func (c *Computer) Ask(question string) (string, error) {
	return c.Model, nil
}

func main() {
	computer := Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	}

	computer.Ask("What is your Model ?")

	answer, err := computer.Ask("What is your Model ?")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Model: %s", answer)
}
