package main

import (
	"go-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	host := "localhost"
	port := "8080"

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Printf("Starting web on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, mux)
	log.Fatal(err)

}
