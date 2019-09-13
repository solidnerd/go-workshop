package main

import "fmt"

type Computer struct {
	Brand string
	Model string
	Price int
}

// TODO: Create a "Ask" method that answers or throw an an error if the question is longer than 20 characters and handle your two return values
func (c *Computer) Ask(question string) (string, error) {
	if len(question) > 20 {
		return "", fmt.Errorf("No")
	}
	return c.Model, nil
}

func main() {
	computer := Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	}

	answer, err := computer.Ask("What is your Model ?")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Model: %s", answer)
}
