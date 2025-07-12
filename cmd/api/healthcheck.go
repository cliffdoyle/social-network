package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information
// about the application status, operating environment and server port its currently running on
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
	}

	err:=app.writeJSON(w,http.StatusOK,data)
	if err !=nil{
	// 	app.logger.Error(err.Error())
	// 	http.Error(w,"The server encountered a problem and not process your request",http.StatusInternalServerError)
	app.serverErrorResponse(w,r,err)
	}
}

func (app *application) errorTest(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		app.methodNotAllowedResponse(w, r)
		return
	}
	err:=fmt.Errorf("simulated server error")
	app.serverErrorResponse(w,r,err)
}
