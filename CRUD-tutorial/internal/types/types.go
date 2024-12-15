package types

// Student struct

type Student struct{
	Id int 
	Name string `json:"name" validate:"required"`
	Age int `json:"age" validate:"required"`
}