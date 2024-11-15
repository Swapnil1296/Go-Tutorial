package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome onboard")
	fmt.Println("Plese rate our service in between 1 to 5")
//  input handling using pkgs provided in go
	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating, ", input);


// converting string to a number 
	numbRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("add one to your ragings : ", numbRating+1)
	}
}