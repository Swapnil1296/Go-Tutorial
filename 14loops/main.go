package main

import "fmt"

func main() {
	fmt.Println("loops in go")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for loop
	fmt.Println(days)

	// for d:=0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }
 
	// for index, day := range days {
	// 	fmt.Printf("Day %d of the week is %s\n", index+1, day)
	// }
	
	somevalue:=1
	for somevalue < 10 {

		// if somevalue==5{
		// 	break
		// }
		if somevalue==2{
			goto statement
		}
		if somevalue ==5{
			somevalue++
			continue
		}
		fmt.Println("value is:",somevalue)
		somevalue++
	}
	statement:
	fmt.Println("This is a statement for goto")
}