package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	// there in nno inheritance in go, but we can achieve it using struct embedding, there is no super or base or parent keyword in go
  swapnil := Employee{name: "Swapnil", empId: 101, role: "Software Engineer", email: "abc@gmail.com", age: 30, status: true}
  fmt.Println("Employee is: ", swapnil)
  fmt.Printf("Employee is : %+v\n", swapnil)
  fmt.Printf("Employee name is: %s\n", swapnil.name)

  //calling the method
  swapnil.showEmployeeStatus()
  swapnil.changeEmployeeEmail()
  fmt.Println("Employee is: ", swapnil)


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
//creating a method with struct to show the employee status

func (e Employee) showEmployeeStatus() {
	fmt.Printf("Employee status is active ? : %t\n", e.status)
}

//manupulating value of the object using method

// this method will not change the original value of the object, it will create a copy of the object and change the value of the copy
func(e Employee) changeEmployeeEmail(){

	e.email = "test@gamil.com"
	fmt.Printf("Employee email is: %s\n", e.email)
}