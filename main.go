package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("layout.html", "genreform.html"))
	err := index.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func handleMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		index := template.Must(template.ParseFiles("layout.html", "genreform.html"))
		err := index.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		r.ParseForm()
		fmt.Println(r.Form)
		for key, value := range r.Form {
			fmt.Printf("key: %s, values: %s", key, value)
		}
		fmt.Fprintln(w, "Music posted successfully")
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/music", handleMusic)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Please navigate to localhost:8080 to view application in browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
