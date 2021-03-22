package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("listen", ":8008", "Override the default listening address")
	flag.Parse()

	log.Println("Binding to " + *addr)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
