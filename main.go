package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("=== UUID Server v1.0 ===")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n", id.String()[:1])
	})

	port := "3000"
	log.Printf("Short UUID server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
