package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod in go lang")
	greeter()
	r:= mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")	
    // http.ListenAndServe(":8080", r)

log.Fatal(http.ListenAndServe(":8080", r))

// http://localhost:8080/  open this link in browser after go run main.go
}



func greeter (){
	fmt.Println("Hello from mymodule")

}



func serveHome(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("<h1>welcome to golang series</h1>"))
}