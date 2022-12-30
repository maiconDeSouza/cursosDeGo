package main

import (
	"log"
	"net/http"
	"text/template"
)

type varHTML struct {
	Patrocionio string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":1992", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	p := varHTML{"JBL"}
	templates.ExecuteTemplate(w, "index.html", p)
}
