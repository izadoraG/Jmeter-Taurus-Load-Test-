package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
