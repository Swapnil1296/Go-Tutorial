package main

import (
	"encoding/json"
	"fmt"
)



type course struct {
	Name string `json:"coursename"`  //providing aliases to the fields
	Price int
	Platform string `json:"website"`
	Password string  `json:"-"`  //this will remove the password field , it will not print
	Tags []string  `json:"tags,omitempty"`  //this will remove the tags field if it is empty
}
func main() {
	fmt.Println("json tutorial")
	// EncodeJson()
	DecodeJson()
}


func EncodeJson() {
	//creating slice of courses

	mycourses := []course{
		{"React JS Tutorial",299,"udemy","12313221",[]string{"web-dev","js"}},
		{"Go Tutorial",299,"udemy","abc123456",nil},
		{"Python Tutorial",399,"udemy","abcdas123456",[]string{"web-dev","js","python"}},

	}
// package this data as JSON data

// finalJson,_ := json.Marshal(mycourses)
finalJson,_ := json.MarshalIndent(mycourses,"","\t")
	fmt.Printf("%s\n",finalJson)

	
}



// how to comsume json data 

func DecodeJson(){

	jsonDataFromWeb := []byte(`
	 {
                "coursename": "Python Tutorial",
                "Price": 399,
                "website": "udemy",
                "tags": ["web-dev","js","python" ]
        }

		
`)
var mycourses course
	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb,&mycourses)
		fmt.Printf("%#v\n",mycourses)
	}else{
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n",k,v,v)
	}
}