package main

import "fmt"

func main() {
	defer fmt.Println("this is going to be printend at the end of the function")
	defer fmt.Println("one")
	defer fmt.Println("two") // this will be printed after the below statement as defer follows LIFO
	fmt.Println("defer in GoLang")

	myDefer() //what will happen here?
	// why the output is 4 3 2 1 0 
	// why the function executed after the  line 9
	//defer is used to delay the execution of a statement until the end of the function or until the surrounding function returns 

}

func myDefer(){

	for i:=0; i<5; i++{
		defer fmt.Println("\n", i)
	}
}