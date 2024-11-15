package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user inputs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating for our pizza: ");

	// comma ok idiom in go  , it is used to check if the value is present or not, if the value is present then it will return true and the value, it is somewhat same as try catch block in other languages
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating, ", input);
	fmt.Printf("Type of input is %T", input)

}