package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"bitbucket.org/gedw99/centri/compute/runner/pkg"
)

func main() {

	pkg.LogStart()

	// Child Servrs
	childUrls := pkg.GetChildWorkerAddresses("work")
	fmt.Printf("children: %v", childUrls)

	http.HandleFunc("/disco", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Disco:")
		fmt.Fprintf(w, "- Children:")
		fmt.Fprintf(w, strings.Join(childUrls, ";"))
	})

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my web server!")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health not implemented")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Http
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
		log.Panicf("NO PORT")
	}
	log.Printf("Starting Server on: %s", port)
	http.ListenAndServe(port, nil)
}
