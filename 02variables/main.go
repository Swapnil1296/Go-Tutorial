package main

import "fmt"


const LoginToken string = "gibberrss"

// variable declared with first cap letters a considered as global variables , it can be access anywher the same file
const loginToken string = "aslkfasdkflsfj" // this is also ok
func main() {
	var variableName string = "Hello, World!"
	fmt.Println(variableName);
	fmt.Printf("variableName is of type %T\n", variableName);

	var isLoggedIng bool = true
	fmt.Println(isLoggedIng);
	fmt.Printf("variableName is of type %T\n", isLoggedIng);

	var AnotherVari uint8 = 255;   // also check for uint16, uint32, uint64
	fmt.Println(AnotherVari);
	fmt.Printf("variableName is of type %T\n", AnotherVari);

	var smallfloat float32 = 255.1565431265365123;
	fmt.Println(smallfloat);
	fmt.Printf("variableName is of type %T\n", smallfloat);

	// default values and some aliases

	var anotherVarible int  // default value will be 0;
    fmt.Println(anotherVarible);
	fmt.Printf("variableName is of type %T\n", anotherVarible);

   // implicit way of declaring variables
    var website = "swapnil" // you can not reassign the value to an int. lexer is deciding the type here
	fmt.Println(website);

	// no var style
	numberofUsers := 3000000   // the := (volorus operator) is allowed to use only in method , it will throw an error if you use it outside the method.
	fmt.Println(numberofUsers);


	fmt.Println(LoginToken);
	fmt.Printf("variableName is of type %T\n", LoginToken);
	fmt.Println(loginToken);
	fmt.Printf("variableName is of type %T\n", loginToken);
	
}


//    fmt.Println(LoginToken)
// this will throw an error as it is not allowed to use fmt.Println() outside the method
// The error you're seeing is because the fmt.Println(LoginToken) statement is outside of the function body, and in Go, non-declaration statements (like function calls) must be inside a function. To fix this, you need to move that fmt.Println(LoginToken) inside the main() function.
// In Go, you cannot place executable code (like fmt.Println()) outside of a function. If you want to print the LoginToken constant outside the main() function or other functions, Go does not allow this directly.
// Go provides a special init() function that runs automatically before the main() function, and it can contain code to print values. This way, you can print your LoginToken without placing the code directly inside the main() function.
func init() {
	// Code in the init function runs before main()
	fmt.Println(LoginToken)
	fmt.Println(loginToken)
}

