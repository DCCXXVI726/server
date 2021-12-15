package main

import "log"

func main() {
	server := NewServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
