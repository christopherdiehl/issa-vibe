package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl = make(map[string]*template.Template)

func parseTemplates() {
	tmpl["getIndex"] = template.Must(template.ParseFiles("layout.html", "genreform.html"))
	tmpl["postIndex"] = template.Must(template.ParseFiles("layout.html", "songPlayer.html"))
}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tmpl["getIndex"].ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		genre := r.FormValue("genre")
		fmt.Println(genre)
		err := tmpl["postIndex"].ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	parseTemplates()
	http.HandleFunc("/", handleRoot)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Please navigate to localhost:8080 to view application in browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
