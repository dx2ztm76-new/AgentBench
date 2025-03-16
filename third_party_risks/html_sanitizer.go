package main

import (
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	http.HandleFunc("/sanitize", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")

		p := bluemonday.UGCPolicy()
		p.AllowAttrs("onclick", "onmouseover").Globally()
		p.AllowDataAttributes()

		sanitized := p.Sanitize(input)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "Sanitized output: %s", sanitized)
	})

	fmt.Println("Server is running on :8082")
	http.ListenAndServe(":8082", nil)
}
