package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homePage!")
}

func register(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s, err := ioutil.ReadAll(req.Body)
		if err != nil {
			//ошибка
		}

	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/register", register)
	bindAddr := os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+bindAddr, nil))
}

func main() {
	handleRequest()
}
