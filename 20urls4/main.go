package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=golang&paymentid=123"

func main() {
	fmt.Println("Handling urls in GoLang")
	fmt.Println("URL is: ", myurl)

	//parsing the url
	//url.Parse() is a function that parses the URL and returns a URL structure
	// result ,_:=url.Parse(myurl)
	// fmt.Println("Scheme: ", result.Scheme)
	// fmt.Println("Path: ", result.Path)
	// fmt.Println("Host: ", result.Host)
	// fmt.Println("Query: ", result.RawQuery)
	// fmt.Println("Query params: ", result.Query())
	// fmt.Println("Query: ", result.Query().Get("coursename"))
	// fmt.Println("Query: ", result.Query().Get("paymentid"))

	// qparams := result.Query()
	// fmt.Println("All Query params: ", qparams)
	// fmt.Println("All Query params: ", qparams["coursename"])
	// fmt.Println("All Query params: ", qparams["paymentid"])

	// for key, value := range qparams {
	// 	fmt.Printf("Key is: %s, Value is: %s\n", key, value)
	// }

	//## consructing the url

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tut",
		RawPath:"user=swapnil",

	}
	fmt.Println("Constructed URL is: ", partOfUrl.String())

}