package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// Create a new random source and rand instance based on the current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
fmt.Println(r)
	// Define the length of the random ID
	length := 10
	id := make([]byte, length)
	fmt.Println(id)
	// Generate the random ID
	for i := range id {
		id[i] = letters[r.Intn(len(letters))]
		fmt.Println(letters[r.Intn(len(letters))])
	}

	// Print the generated random ID
	fmt.Println("Random ID:", string(id))

	
}
