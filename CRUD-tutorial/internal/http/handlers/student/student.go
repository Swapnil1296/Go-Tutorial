package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Swapnil1296/crud-tutorial/internal/types"
	response "github.com/Swapnil1296/crud-tutorial/internal/utils"
	"github.com/go-playground/validator/v10"
)

func StudentHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		var student types.Student
		//decode the request body if it is empty then return error
		err :=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
		response.WriteJson(w,http.StatusBadGateway,response.GeneralError(fmt.Errorf("empty request body")))
		return
		}
		// if error is not empty body then
if err != nil{
	response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
	return
}
//validate the input 
if err:=validator.New().Struct(student); err != nil{
	// typecast the error to validation error
	validateErrs :=err.(validator.ValidationErrors)
	//return the error
	response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
	return
}  


   
response.WriteJson(w,http.StatusCreated,map[string]string{"success":"ok"})
}
}