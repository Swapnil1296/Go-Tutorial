package main

import "fmt"

func main() {
	fmt.Println("Pointers in GoLang")
	var prt *int   // creating a pointer variable
	fmt.Println("value of pointer: ", prt);

	mynumber :=23;

	var ptrr = &mynumber; // assigning address of mynumber to ptrr  (when pointer is referencing to a variable)
	fmt.Println("value of pointer: ", ptrr);
	fmt.Println("value of pointer: ", *ptrr);
	*ptrr = *ptrr + 5;
	fmt.Println("value of pointer now: ", *ptrr);

}