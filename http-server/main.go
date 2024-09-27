package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello`))
	})
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "HTTP method is wrong", http.StatusNotFound)
			return
		}

		search := r.URL.Query().Get("search")
		fmt.Println("search", search)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo":"bar"}`))
	})
	http.ListenAndServe(":4444", nil)
}
