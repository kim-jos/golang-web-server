package main

import (
	"net/http"
	myApp "web-server/my-app"
)

func main() {
	http.ListenAndServe(":3000", myApp.NewHttpHandler())
}
