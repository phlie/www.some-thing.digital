package main

import (
	//"fmt"
	"log"
"os"
	"html/template"
	"regexp"
	"net/http"
)

var validPath = regexp.MustCompile(``)

func loadIndexPage () ([]byte, error) {
	filename := "index.html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 1 {
		http.NotFound(w, r)
		return
	}
	//	m := validPath.FindString(r.URL.Path[0])
	page, err := loadIndexPage()
	if err != nil {
		return
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	t.Execute(w, page)
}

func main() {
	m := http.NewServeMux()
	m.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("css"))))
	m.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs"))))
	m.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8080", m))
}
