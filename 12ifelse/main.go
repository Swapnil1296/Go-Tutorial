package main

import "fmt"

func main() {
	fmt.Println("If Else in GoLang")

	count := 11
    result := ""
	if count < 10 {
		result = "Less than 10"
	}else if count == 10 {
		result = "Equal to 10"
	}else {
		result = "Greater than 10"
	}
	fmt.Println("Result is: ", result)

	// if else with initialization
	if num := 10; num < 10 {
		fmt.Println("Less than 10")
	}else if num == 10 {
		fmt.Println("Equal to 10")
	}else {
		fmt.Println("Greater than 10")
	}
// if else with multiple conditions
	if 5%4 == 0 {
		fmt.Println("5 is divisible by 4")
	}else {
		fmt.Println("5 is not divisible by 4")
	}
}