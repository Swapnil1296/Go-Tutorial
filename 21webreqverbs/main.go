package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	// PerformGetRequest()
	// PerformPostRequestJsonData()
	PerformPostRequestFormData()
}


func PerformGetRequest() {
	fmt.Println("Performing GET Request")
	const url = "http://localhost:8000/get"

	response,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
		
	}
	defer response.Body.Close()
	fmt.Println("Status Code:",response.Status)
	fmt.Println("content length:",response.ContentLength)

	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))

	// using content builder to read the response
	 var contentBuilder strings.Builder
	 bytecount,_ := contentBuilder.Write(content)
     fmt.Println("Content Builder:",bytecount)
	 fmt.Println("Content Builder content:",contentBuilder.String())
}

func PerformPostRequestJsonData() {

	const url = "http://localhost:8000/post"

	// json data
	requestBody := strings.NewReader(`{"name":"John Doe","age":25}`)

	response,err:=http.Post(url,"application/json",requestBody)
	if err!=nil{
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:",response.Status)
	fmt.Println("content length:",response.ContentLength)
	fmt.Println("content type:",response.Header.Get("Content-Type"))
	fmt.Println("content length:",response.Header.Get("Content-Length"))
	fmt.Println("response:",response)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostRequestFormData() {
	
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}

	data.Add("name","John Doe")
	data.Add("age","25")
	data.Add("email","john@gmail.com")

	res,_:=http.PostForm(myurl,data)

	
	fmt.Println("Status Code:",res.Status)
	fmt.Println("content length:",res.ContentLength)
	fmt.Println("content type:",res.Header.Get("Content-Type"))	
	fmt.Println("respnonse",res)

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}