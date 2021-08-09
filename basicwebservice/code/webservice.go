package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed test.txt
var textfile []byte

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	return &Page{Title: title, Body: textfile}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	fmt.Println(string(textfile))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
