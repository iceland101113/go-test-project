package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Println("Starting server....")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/greeting", Greeting)
	http.Serve(listener, nil)
}


func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helloe Gopher")
}