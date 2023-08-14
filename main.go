package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/products.html",
		"templates/contact.html",
		"templates/public/header.html",
		"templates/public/footer.html",
		"templates/public/css.html",
		"templates/public/scripts.html",
	))
	// static file server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "contact.html", nil)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "products.html", nil)
	})

	fmt.Println("Server is running on port 8050")

	// start server
	http.ListenAndServe(":8050", nil)
}
