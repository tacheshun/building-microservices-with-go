package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world")
	})

	log.Fatal(http.ListenAndServe(":9090", nil))
}
