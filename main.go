package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("layout.html", "index.html"))
	err := index.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", handleRoot)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Please navigate to localhost:8080 to view application in browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
