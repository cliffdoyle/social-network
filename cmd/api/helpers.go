package main

import (
	"encoding/json"
	"net/http"
)

//Define a writeJSON() helper for sending responses, takes
//the destination http.ResponseWriter, the HTTP status code to send and the 
//data to encode to JSON

func (app *application) writeJSON(w http.ResponseWriter, status int, data any)error{
	//Encode the data to JSON, returning the error if there is one
	js,err:=json.MarshalIndent(data,"","\t")
	if err !=nil{
		return  err
	}

	//Append a newline to make it easier to view in terminal applications
	js = append(js, '\n')

	//Add the "Content-Type: application/json" header, then
	//write the status code and JSON response.
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write(js)

	return  nil
}
