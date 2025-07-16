package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler)http.Handler{
	//Initialize a new rate limiter which allows an average of 2 requests per second,
	//With a maximum of 4 requests in a single 'burst'
	
	limiter:=rate.NewLimiter(2,4)
	
	//The function we are returning is a closure, which 'closes over' the limiter variable.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//Call limiter.Allow() to see if the request is permitted, and if it's not,
	//then call rateLimitExceededResponse() helper to return a 429 too many
	//Requests response (we will create this helper in a minute)
	if !limiter.Allow(){
		app.rateLimitExceededResponse(w,r)
		return
	}

	next.ServeHTTP(w,r)

	})
}


