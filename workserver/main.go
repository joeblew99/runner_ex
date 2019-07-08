package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joeblew99/runner_ex/pkg"
)

var (
	duration = 2
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

	// Add 5 second delay to simulate long startup time
	fmt.Printf("Sleeping for %v seconds\n", duration)
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Slept for %v seconds\n", duration)

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my worker!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Server
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, nil)
}
