package main

import (
	"fmt"
	"net/http"
)


//Declare a handler which writes a plain-text response with information
//about the application status, operating environment and server port its currently running on
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"status:available")
	fmt.Fprintf(w,"environment:%s\n",app.config.env)
	fmt.Fprintf(w,"server port:%d\n",app.config.port)
}