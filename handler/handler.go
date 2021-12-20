package handler

import (
	"go-web/entity"
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

	data := map[string]interface{}{
		"title":   "Learning Go Web",
		"content": []string{"index", "html"},
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	data := entity.Product{Id: 1, Name: "Car", Price: 1200000000, Stock: 3}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"title":   "Detail Product",
		"content": idNum,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
