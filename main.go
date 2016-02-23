package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting ...")
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World, %s", req.URL.Path[1:])
}
