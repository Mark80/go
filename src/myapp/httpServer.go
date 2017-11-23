package main

import (
	"myapp/server"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", &server.MyMx{})
}
