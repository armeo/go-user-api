package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)

	fmt.Println("Server starting ...")
	http.ListenAndServe(":8000", nil)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(map[string]string{"message": "Hello"})

	res.Write(msg)
}
