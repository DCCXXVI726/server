package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homePage!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	bindAddr :=  os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8081"
	}
	log.Fatal(http.ListenAndServe(":" + bindAddr, nil))
}

func main() {
	handleRequest()
}
