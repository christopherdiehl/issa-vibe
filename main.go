package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl = make(map[string]*template.Template)
var playlists = map[string]string{
	"Punk":    "37i9dQZF1DXa9wYJr1oMFq",
	"Jazz":    "37i9dQZF1DWVqfgj8NZEp1",
	"Blues":   "37i9dQZF1DXaiAJKcabR16",
	"Pop":     "37i9dQZF1DXcBWIGoYBM5M",
	"Rock":    "37i9dQZF1DXcF6B6QPhFDv",
	"Hip Hop": "37i9dQZF1DX0XUsuxWHRQd",
	"Folk":    "37i9dQZF1DWYV7OOaGhoH0",
	"Indie":   "37i9dQZF1DX2Nc3B70tvx0",
}

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
		var url strings.Builder
		id := playlists[genre]
		url.WriteString("https://open.spotify.com/embed/track/")
		url.WriteString(id)
		url.WriteString("FOO")
		fmt.Println(url.String())
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
