package main

import (
	"fmt"
)

func main() {

fmt.Println("Maps in GoLang")

	//creating a map
 var fruitRank = make(map[string]string)

 fruitRank["mg"] = "Mango"
 fruitRank["ap"] = "Apple"
 fruitRank["bn"] = "Banana"
 fruitRank["or"] = "Orange"
 fmt.Println("Fruit Rank is: ", fruitRank)
 fmt.Println("ap: ", fruitRank["ap"])
 delete(fruitRank, "mg")
 fmt.Println("Fruit Rank is: ", fruitRank)

 //loops in maps

 for key , value := range fruitRank {
	 fmt.Printf("Fruit Rank for %s is %s\n", key, value)
 }
}