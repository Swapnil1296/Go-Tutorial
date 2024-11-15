package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	// there in nno inheritance in go, but we can achieve it using struct embedding, there is no super or base or parent keyword in go
  swapnil := Employee{name: "Swapnil", empId: 101, role: "Software Engineer", email: "abc@gmail.com", age: 30, status: true}
  fmt.Println("Employee is: ", swapnil)
  fmt.Printf("Employee is : %+v\n", swapnil)
  fmt.Printf("Employee name is: %s\n", swapnil.name)

}

//creating a struct
type Employee struct {
	name string
	empId int
	role string
	email string
	age int
	status bool
}