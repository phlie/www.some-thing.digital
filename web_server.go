package main

// All the packages and only the packages that the web server uses.
import (
	//"fmt"
	"log"
	"os"
	"html/template"
	"regexp"
	"net/http"
)

// The only valid web address after our base address that is valid for the index page
var validPath = regexp.MustCompile(``)

// Gets the index page and returns it with it in byte array form, if there is an error it will return that error
func loadIndexPage () ([]byte, error) {
	filename := "index.html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// The function responsible for checking to see if the path in the browser bar is valid
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// If the path is not exactly just a forward slash, ie. www.some-thing.digital/ then return
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// If there are any GET Queries, also return a not found
	if len(r.URL.Query()) != 0 {
		http.NotFound(w, r)
		return
	}
	// If its address is valid load the page
	page, err := loadIndexPage()
	// If there is an error, return
	if err != nil {
		return
	}
	// Parse the file throwing an error if it can't be done.
	t, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	// Finally, display the page
	t.Execute(w, page)
}

// The Main function, the first one called in the program only the index.html, imgs in /imgs/ and files in the /styles/ directories are at all accessible to viewers of our page.
func main() {
	// Create a new serve request
	m := http.NewServeMux()
	// Handles the Style sheet in the css directory only allowing that one file to be accessed.
	m.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("css"))))
	// Allows only images in the imgs directory to be accessed.
	m.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs"))))
	// Finally, the Index Page
	m.HandleFunc("/", mainHandler)
	// If there is a fatal error while serving, log it to the console
	// For the actual release it should listen on port 80, or ":80"
	log.Fatal(http.ListenAndServe(":8080", m))
}
