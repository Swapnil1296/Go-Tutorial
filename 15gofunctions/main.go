package main

import "fmt"

func main(){
	fmt.Println("Functions in GoLang")
	greeter()
	greeterTwo()
	// ***you can not decalre a function inside a function, it will give an error
	result := adder(2,6)
	fmt.Println("Result is:",result)

	proRes := proAdder(1,2,3,4,5,6,7,8,9)
	fmt.Println("Pro Result is:",proRes)

	//to get the multiple return values from a function you need to use the below syntax, this is just like comma ok idiom , if you dont want to use the return value you can use the _ operator
	proResMulti, str := proAdderMulti(1,2,3,4,5)
	fmt.Println("Pro Result is:",proResMulti)
	fmt.Println("String is:",str)
}
//function signature are the function name and the parameters, the return type is not part of the signature in go you need declare the return type in this case int is the return type
func adder(value1 int, value2 int) int{
	return value1 + value2

}
// you can also pass multiple values to a function using the ... operator
func proAdder(values ...int) int{
	total := 0
	for _, value := range values{
		total += value
	}
	return total
}
// you can return multiple types of values from a function, in the below example we are returning an int and a string
func proAdderMulti(values ...int)( int,string){
	total := 0
	for _, value := range values{
		total += value
	}
	return total , "This is a string"
}
func greeterTwo(){
	fmt.Println("Hello Golang")
}
func greeter(){
	fmt.Println("Hello World")
}