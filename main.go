package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Binding to :8008")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Fatal(http.ListenAndServe(":8008", nil))
}
