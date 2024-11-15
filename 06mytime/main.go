package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("this is my time : ")
	presentTime := time.Now();
    fmt.Println("current Time: ", presentTime)
	fmt.Println("current Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))	// Monday is mandatory go get exact day.
  createdDate := time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC)
  fmt.Println("Created Date: ", createdDate);
  fmt.Println("Created Date: ", createdDate.Format("02-01-2006 15:04:05 Monday"))	// Monday is mandatory go get exact day.


}