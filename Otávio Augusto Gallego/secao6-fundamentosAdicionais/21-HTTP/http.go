package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":1992", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!!!"))
}
