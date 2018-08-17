package main

import(
	"log"
	"net/http"
	"fmt"
	"html/template"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("layout.html","index.html"))
	err := index.ExecuteTemplate(w, "index.html",nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main(){
	http.HandleFunc("/", handleRoot)
	fmt.Println("Please navigate to localhost:8080 to view application in browser")
	log.Fatal(http.ListenAndServe(":8080",nil))
}