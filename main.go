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
	mux.HandleFunc("/products", handler.ProductsHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/get-post", handler.GetPost)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting web on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, mux)
	log.Fatal(err)

}
