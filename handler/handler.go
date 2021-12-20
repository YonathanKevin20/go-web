package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Learning Go Web",
		"content": []int{1, 2},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product: %v", id)
}
