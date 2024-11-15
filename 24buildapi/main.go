package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Model for course file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`	

}
type Author struct{
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Website string `json:"website"`
}
// fake DB
var courses []Course

//middleware/helper function

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0
	return c.CourseName ==""
	
}
func main() {
	fmt.Println("Hello, World!")

	// define the router
	r := mux.NewRouter()
	// seeding 
	courses = append(courses,Course{CourseId:"1",CourseName:"Python",CoursePrice:1000,Author:&Author{FullName:"John Doe",Email:"test@gmail.com",Website:"www.test.com"}} )
	courses =append(courses,Course{CourseId:"2",CourseName:"Reac JS",CoursePrice:1000,Author:&Author{FullName:"Sha Cloe",Email:"test@gmail.com",Website:"www.test.com"}} )

	// define the routes
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllcourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",creatOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")


	//hadle error
	log.Fatal(http.ListenAndServe(":8080",r))
} 

//contorllers for CRUD operations

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}


func getAllcourses(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
	
}


func getOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	//grab id from the request 
	// id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	// loop through the courses and find the course with the id
	for _,course := range courses{
		if course.CourseId == params["id"]{
			// sends the course as a response
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	
}

func creatOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
// what if body is empty

if r.Body == nil{
	// http.Error(w,"Please send some data",http.StatusBadRequest)

	json.NewEncoder(w).Encode("Please send some data")	
	return	
}
// What if there in no data in {}

var course Course 

_=json.NewDecoder(r.Body).Decode(&course)
if course.IsEmpty(){
	json.NewEncoder(w).Encode("No data found inside JSON")
	return
}
// generate a unique id, convert to string
// append course into courses
// Create a new random source and rand instance based on the current time
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
source := rand.NewSource(time.Now().UnixNano())
result := rand.New(source)

// Define the length of the random ID
length := 10
id := make([]byte, length)

// Generate the random ID
for i := range id {
	id[i] = letters[result.Intn(len(letters))]
}
course.CourseId=string(id)

json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//grab id from the request 
	// id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	// loop through the courses and find the course with the id
	for index,course := range courses{
		if course.CourseId == params["id"]{
			// sends the course as a response
			courses = append(courses[:index],courses[index+1:]...)
			var course Course 
			_=json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}


func deleteOneCourse (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//grab id from the request 
	// id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	// loop through the courses and find the course with the id
	for index,course := range courses{
		if course.CourseId == params["id"]{
			// sends the course as a response
			courses = append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")

}
//seeding : adding some fake data to DB 