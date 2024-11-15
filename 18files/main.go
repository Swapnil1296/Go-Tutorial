package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("handeling files in GoLang")

	

	content := "Hello, this is a file handling in GoLang"
	//creating a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	length, err := io.WriteString(file,content)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println("length of the file is: ", length)
	defer file.Close()

	readFile("test.txt")

}

func readFile (filename string){

	databytes,err:=os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Data from the file is: ", string(databytes))
}

// we can create this type of function to check the error and print the error instead of writing the same code again and again
func checkNilError(err error){
	if err != nil {
		fmt.Println(err)
		return
	}
}