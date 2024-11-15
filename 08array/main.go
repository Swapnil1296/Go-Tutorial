package main

import "fmt"

func main() {
	fmt.Println("Arrays in GoLang");

	var fruitList [4]string
	fruitList[0] = "tomato"
	fruitList[1]="papaya"
	fruitList[3]="apple"
	
	fmt.Println("Fruit List is: ", fruitList);
	fmt.Println("Fruit List length is: ", len(fruitList));

  var vegList = [3]string{"potato", "onion", "brinjal"}   //creating an array & assiging values immediately;
  fmt.Println("Veg List is: ", vegList);
  fmt.Println("Veg List length is: ", len(vegList));


}