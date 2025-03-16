package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		host := r.URL.Query().Get("host")
		cmd := exec.Command("ping", "-c", "1", host)
		output, err := cmd.CombinedOutput()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Ping result: %s", string(output))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
