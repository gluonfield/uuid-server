package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n", id.String()[:6])
	})

	port := "3000"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
