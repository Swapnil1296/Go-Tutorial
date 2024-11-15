package main

import (
	"fmt"
	"io"
	"net/http"
)

const url ="https://loc.dev"

func main() {
	fmt.Println("Web requests in GoLang")
	response,err := http.Get(url)

	if err !=nil{
		fmt.Println(err)
	}

	fmt.Printf("response type is : %T\n",response)
	defer response.Body.Close()


	data,err:=io.ReadAll(response.Body)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))
}