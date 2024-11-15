package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in GoLang")
	rand.New(rand.NewSource(time.Now().UnixNano()));
	// Generate a random number between 1 and 5
	randomNumber := rand.Intn(6) + 1
	fmt.Println("Random Number: ", randomNumber)
	switch randomNumber {
	case 1: fmt.Println("randomNumber is one")
	case 2: fmt.Println("randomNumber is two")

	case 3: fmt.Println("randomNumber is three")
	case 4: fmt.Println("randomNumber is four")
	fallthrough // fallthrough is used to execute the next case block as well
	case 5: fmt.Println("randomNumber is five")
	fallthrough
	default: fmt.Println("randomNumber is not between 1 and 5")
	}
}