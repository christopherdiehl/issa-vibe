package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl = make(map[string]*template.Template)
var genreTrackMap = map[string][]string{
	"Punk":    []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Jazz":    []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Blues":   []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Pop":     []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Rock":    []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Hip Hop": []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Folk":    []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Indie":   []string{"2D7tauy2bntJnJQ2C4rO4x"},
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
		var url strings.Builder
		id := genreTrackMap[genre]
		// for track use https://open.spotify.com/embed/track/
		url.WriteString("https://open.spotify.com/embed/track/")
		url.WriteString(id)
		fmt.Println(url.String())
		err := tmpl["postIndex"].ExecuteTemplate(w, "base", &struct{ URL string }{url.String()})
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
