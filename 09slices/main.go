package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slices in GoLang");

	var fruitList = []string{"papaya", "apple","graps"}

fmt.Printf("Type of fruitList is %T\n", fruitList);
fmt.Println("Fruit List is: ", fruitList);
fruitList = append(fruitList, "banana");
fmt.Println("Fruit List is: ", fruitList);


fruitList = append(fruitList[1:]); //removing first element
fmt.Println("removing first element: ", fruitList);
fruitList = append(fruitList[:len(fruitList)-1]); //removing last element
fmt.Println("removing last element: ", fruitList);
fruitList = append(fruitList[1:2])
fmt.Println("removing between range: ", fruitList);


highScores := make([]int, 4)

highScores[0] = 234
highScores[1] = 945
highScores[2] = 465
highScores[3] = 867
// highScores[4] = 869 //this will throw an error as the size of the array is 4 & we are trying to assign 5th element
fmt.Println("High Scores length: ", len(highScores))
fmt.Println("High Scores: ", highScores)
highScores = append(highScores, 555, 666, 777)	//this will create a new array with double the size & copy the old values to new array & then append the new values to the new array & assign it to the old array
fmt.Println("High Scores length after adding more values: ", len(highScores))

fmt.Println("High Scores after adding more values: ", highScores)

sort.Ints(highScores); //sorting the array, this will sort the array in ascending order, if we want to sort in descending order we can use sort.Sort(sort.Reverse(sort.IntSlice(highScores)))
fmt.Println("High Scores after sorting: ", highScores)

// how to remove an element from a slice based on index

var courses = []string{"react", "angular", "vue", "svelte", "jvm", "springboot", "kotlin", "go"}
fmt.Println("Courses: ", courses);

var index int = 2;
// courses = append(courses[:index], courses[index+1:]); // this will throw an error
courses = append(courses[:index], courses[index+1:]...); // this will remove the element at index 2
fmt.Println("Courses after removing an element: ", courses);	
}