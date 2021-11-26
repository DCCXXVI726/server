package src

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homePage!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	bindAddr :=  os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8080"
	}
	http.ListenAndServe(":" + bindAddr, nil)
}

func main() {
	handleRequest()
}