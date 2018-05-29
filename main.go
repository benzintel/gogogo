package main

import (
	user "./controller"
	"net/http"
)

func main() {

	http.HandleFunc("/", user.SayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}